package kube

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/chaos/internal/awscli"
	"github.com/Ningendo7/terraform-multi-region-platform/tools/chaos/internal/safety"
)

const (
	cordonedByAnnotation   = "chaos.tools/cordoned-by"
	restoreAfterAnnotation = "chaos.tools/restore-after"
	nodeFailureExperiment  = "node-failure"

	// Refuse to take out more than half the cluster by default — a bug
	// in --az matching (or a mistyped zone) taking out every node should
	// be structurally hard, not just something to be careful about.
	defaultBlastRadiusCap = 0.5
)

// selectByCount deterministically samples count items from candidates
// using seed — pulled out as its own pure function so the sampling
// logic is testable without shelling out to kubectl. count <= 0 or
// count >= len(candidates) means "use everything," unchanged.
func selectByCount(candidates []string, count int, seed int64) []string {
	if count <= 0 || count >= len(candidates) {
		return candidates
	}

	rng := rand.New(rand.NewSource(seed))

	shuffled := make([]string, len(candidates))
	copy(shuffled, candidates)
	rng.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })

	return shuffled[:count]
}

// checkBlastRadius refuses a target set larger than the cap unless force
// is set — pulled out as its own pure function so the 50% guardrail is
// itself unit-tested, not just exercised implicitly through NodeFailure.
func checkBlastRadius(targetCount, clusterSize int, force bool) error {
	if force {
		return nil
	}

	ratio := float64(targetCount) / float64(clusterSize)
	if ratio > defaultBlastRadiusCap {
		return fmt.Errorf("refusing to target %d of %d nodes (%.0f%% of the cluster, over the %.0f%% cap) — pass --force to override",
			targetCount, clusterSize, ratio*100, defaultBlastRadiusCap*100)
	}

	return nil
}

// NodeFailure cordons and evicts a set of real nodes, selected either by
// count (N random nodes cluster-wide) or az (every node — or up to count
// of them — in one availability zone). This one primitive covers both
// "kill some nodes" and "simulate losing a zone": az-failure is just this
// with az set and count left at its zone-wide default of "all of them."
//
// Cordoned nodes restore automatically after window, or earlier via a
// separate `chaos restore` call — see waitAndRestore.
func NodeFailure(ctx context.Context, kubeContext, awsRegion, az string, count int, window time.Duration, force, dryRun, yes bool, seed int64) error {
	if err := safety.RequireConfirmation(dryRun, yes); err != nil {
		return err
	}

	allNodes, err := getNodes(ctx, kubeContext)
	if err != nil {
		return err
	}

	var candidates []string
	if az != "" {
		cidrs, err := awscli.CIDRsForAZ(ctx, awsRegion, az)
		if err != nil {
			return err
		}
		candidates, err = nodeNamesInCIDRs(allNodes, cidrs)
		if err != nil {
			return err
		}
	} else {
		for _, n := range allNodes {
			candidates = append(candidates, n.Metadata.Name)
		}
	}

	if len(candidates) == 0 {
		if az != "" {
			return fmt.Errorf("no nodes found in %s", az)
		}
		return fmt.Errorf("no nodes found in cluster")
	}

		if seed == 0 {
		seed = time.Now().UnixNano()
	}
	candidates = selectByCount(candidates, count, seed)

	if err := checkBlastRadius(len(candidates), len(allNodes), force); err != nil {
		return err
	}

	if dryRun {
		fmt.Printf("action=dry-run targets=%v az=%q count=%d cluster_size=%d context=%s\n",
			candidates, az, len(candidates), len(allNodes), kubeContext)
		return nil
	}

	restoreAfter := time.Now().Add(window)

	for _, node := range candidates {
		if err := cordonNode(ctx, kubeContext, node); err != nil {
			return fmt.Errorf("cordoning %s: %w", node, err)
		}
		if err := annotateNode(ctx, kubeContext, node, cordonedByAnnotation, nodeFailureExperiment); err != nil {
			return fmt.Errorf("annotating %s: %w", node, err)
		}
		if err := annotateNode(ctx, kubeContext, node, restoreAfterAnnotation, restoreAfter.Format(time.RFC3339)); err != nil {
			return fmt.Errorf("annotating %s: %w", node, err)
		}
		if err := evictPods(ctx, kubeContext, node); err != nil {
			return fmt.Errorf("evicting pods on %s: %w", node, err)
		}
	}

	fmt.Printf("action=cordoned targets=%v az=%q restore_after=%s context=%s\n",
		candidates, az, restoreAfter.Format(time.RFC3339), kubeContext)
	fmt.Printf("waiting up to %s — run `chaos restore --context %s` from another terminal to restore early\n", window, kubeContext)

	return waitAndRestore(ctx, kubeContext, restoreAfter)
}

// waitAndRestore blocks until either the deadline passes or an external
// `chaos restore` call has already un-annotated these nodes, then
// restores anything still cordoned by this experiment. Node annotations
// are the shared state — no separate daemon or local file needed for a
// second process to signal "restored early."
func waitAndRestore(ctx context.Context, kubeContext string, deadline time.Time) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		nodes, err := getNodes(ctx, kubeContext)
		if err != nil {
			return err
		}

		if len(nodeNameWithAnnotation(nodes, cordonedByAnnotation, nodeFailureExperiment)) == 0 {
			fmt.Println("action=already-restored — an external `chaos restore` beat the window")
			return nil
		}

		if time.Now().After(deadline) {
			return Restore(ctx, kubeContext)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}

// Restore uncordons every node this tool has cordoned and not yet
// restored — the safety net for "the operator forgot," and also what a
// manual `chaos restore` call does to end an experiment early.
func Restore(ctx context.Context, kubeContext string) error {
	nodes, err := getNodes(ctx, kubeContext)
	if err != nil {
		return err
	}

	targets := nodeNameWithAnnotation(nodes, cordonedByAnnotation, nodeFailureExperiment)
	if len(targets) == 0 {
		fmt.Println("action=restore nothing to restore")
		return nil
	}

	for _, node := range targets {
		if err := uncordonNode(ctx, kubeContext, node); err != nil {
			return fmt.Errorf("uncordoning %s: %w", node, err)
		}
		if err := removeNodeAnnotation(ctx, kubeContext, node, cordonedByAnnotation); err != nil {
			return fmt.Errorf("removing annotation on %s: %w", node, err)
		}
		if err := removeNodeAnnotation(ctx, kubeContext, node, restoreAfterAnnotation); err != nil {
			return fmt.Errorf("removing annotation on %s: %w", node, err)
		}
	}

	fmt.Printf("action=restored targets=%v context=%s\n", targets, kubeContext)
	return nil
}
