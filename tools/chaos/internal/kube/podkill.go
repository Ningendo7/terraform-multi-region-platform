// tools/chaos/internal/kube/podkill.go
package kube

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// PodKill deletes one pod matching label within namespace, on the cluster
// identified by kubeContext.
//
// seed makes target selection reproducible — pass 0 to seed from the
// current time. Every run logs one structured line naming exactly what it
// selected and why, dry-run or not, so there's always an audit trail.
//
// A real (non-dry-run) delete requires yes=true — refusing to delete
// without either --dry-run or explicit --yes is deliberate: this tool's
// whole job is destructive, so the destructive path should never be the
// accidental default.
func PodKill(ctx context.Context, kubeContext, namespace, label string, dryRun, yes bool, seed int64) error {
	if !dryRun && !yes {
		return fmt.Errorf("refusing to delete without --dry-run or --yes")
	}

	out, err := Run(ctx, kubeContext, "get", "pods", "-n", namespace, "-l", label, "-o", "name")
	if err != nil {
		return err
	}

	pods := strings.Fields(out)
	if len(pods) == 0 {
		return fmt.Errorf("no pods matched namespace=%s label=%s", namespace, label)
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rng := rand.New(rand.NewSource(seed))
	target := pods[rng.Intn(len(pods))]

	logLine := func(action string) {
		fmt.Printf("action=%s target=%s namespace=%s label=%q candidates=%d seed=%d context=%s\n",
			action, target, namespace, label, len(pods), seed, kubeContext)
	}

	if dryRun {
		logLine("dry-run")
		return nil
	}

	if _, err := Run(ctx, kubeContext, "delete", target, "-n", namespace); err != nil {
		return err
	}

	logLine("deleted")
	return nil
}
