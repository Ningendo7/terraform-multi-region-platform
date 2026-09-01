package chaoscli

import (
	"context"
	"fmt"
	"time"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/chaos/internal/kube"
)

func Execute(args []string) error {

	if len(args) == 0 {
		Usage()
		return nil
	}

	switch args[0] {

	case "run":

		if len(args) < 2 {
			return fmt.Errorf("run requires an experiment name")
		}

		return runExperiment(args[1], args[2:])

	case "restore":

		opts, err := ParseRestoreFlags(args[1:])
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		return kube.Restore(ctx, opts.Context)

	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runExperiment(name string, args []string) error {

	switch name {

	case "pod-kill":

		opts, err := ParsePodKillFlags(args)
		if err != nil {
			return err
		}

		// 30s covers this experiment's two sequential kubectl calls with
		// real headroom, and guarantees the tool exits rather than hangs
		// against an unreachable cluster.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		return kube.PodKill(ctx, opts.Context, opts.Namespace, opts.Label, opts.DryRun, opts.Yes, opts.Seed)

	case "node-failure":

		opts, err := ParseNodeFailureFlags(args)
		if err != nil {
			return err
		}

		// Unlike pod-kill, this one deliberately blocks for up to
		// --window waiting to auto-restore, so the timeout has to cover
		// the whole window plus real headroom for the cordon/evict
		// calls — not a fixed short duration.
		ctx, cancel := context.WithTimeout(context.Background(), opts.Window+2*time.Minute)
		defer cancel()
		return kube.NodeFailure(ctx, opts.Context, opts.Region, opts.AZ, opts.Count, opts.Window, opts.Force, opts.DryRun, opts.Yes, opts.Seed)

	default:
		return fmt.Errorf("unknown experiment: %s", name)
	}
}
