package chaoscli

import (
	"flag"
	"fmt"
	"io"
	"time"
)

type PodKillOptions struct {
	Context   string
	Namespace string
	Label     string
	DryRun    bool
	Yes       bool
	Seed      int64
}

func ParsePodKillFlags(args []string) (PodKillOptions, error) {

	fs := flag.NewFlagSet("pod-kill", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var opts PodKillOptions
	fs.StringVar(&opts.Context, "context", "", "kubeconfig context of the target cluster (required)")
	fs.StringVar(&opts.Namespace, "namespace", "", "namespace to target (required)")
	fs.StringVar(&opts.Label, "label", "", "label selector for candidate pods, e.g. app=demo-app (required)")
	fs.BoolVar(&opts.DryRun, "dry-run", false, "select a target and log it without deleting anything")
	fs.BoolVar(&opts.Yes, "yes", false, "required to actually delete a pod (omit to require --dry-run)")
	fs.Int64Var(&opts.Seed, "seed", 0, "seed for reproducible target selection (0 = seed from current time)")

	if err := fs.Parse(args); err != nil {
		return PodKillOptions{}, err
	}

	if opts.Context == "" || opts.Namespace == "" || opts.Label == "" {
		return PodKillOptions{}, fmt.Errorf("pod-kill requires --context, --namespace, and --label")
	}

	return opts, nil
}

type NodeFailureOptions struct {
	Context string
	Region  string
	AZ      string
	Count   int
	Window  time.Duration
	Force   bool
	DryRun  bool
	Yes     bool
	Seed    int64
}

func ParseNodeFailureFlags(args []string) (NodeFailureOptions, error) {

	fs := flag.NewFlagSet("node-failure", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var opts NodeFailureOptions
	fs.StringVar(&opts.Context, "context", "", "kubeconfig context of the target cluster (required)")
	fs.StringVar(&opts.Region, "region", "", "AWS region the cluster lives in (required if --az is set)")
	fs.StringVar(&opts.AZ, "az", "", "restrict targets to this availability zone, e.g. us-east-1a")
	fs.IntVar(&opts.Count, "count", 0, "number of nodes to target (required if --az is not set; with --az, 0 means every node in that AZ)")
	fs.DurationVar(&opts.Window, "window", 5*time.Minute, "how long targeted nodes stay cordoned before auto-restoring")
	fs.BoolVar(&opts.Force, "force", false, "override the default 50% blast-radius cap")
	fs.BoolVar(&opts.DryRun, "dry-run", false, "select targets and log them without touching anything")
	fs.BoolVar(&opts.Yes, "yes", false, "required to actually run this experiment (omit to require --dry-run)")
	fs.Int64Var(&opts.Seed, "seed", 0, "seed for reproducible target selection when --count is used (0 = seed from current time)")

	if err := fs.Parse(args); err != nil {
		return NodeFailureOptions{}, err
	}

	if opts.Context == "" {
		return NodeFailureOptions{}, fmt.Errorf("node-failure requires --context")
	}

	if opts.AZ == "" && opts.Count <= 0 {
		return NodeFailureOptions{}, fmt.Errorf("node-failure requires --count when --az is not set")
	}

	if opts.AZ != "" && opts.Region == "" {
		return NodeFailureOptions{}, fmt.Errorf("node-failure requires --region when --az is set")
	}

	return opts, nil
}

type RestoreOptions struct {
	Context string
}

func ParseRestoreFlags(args []string) (RestoreOptions, error) {

	fs := flag.NewFlagSet("restore", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var opts RestoreOptions
	fs.StringVar(&opts.Context, "context", "", "kubeconfig context of the target cluster (required)")

	if err := fs.Parse(args); err != nil {
		return RestoreOptions{}, err
	}

	if opts.Context == "" {
		return RestoreOptions{}, fmt.Errorf("restore requires --context")
	}

	return opts, nil
}
