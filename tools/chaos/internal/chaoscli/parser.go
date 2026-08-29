// tools/chaos/internal/chaoscli/parser.go
package chaoscli

import (
	"flag"
	"fmt"
	"io"
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
	fs.SetOutput(io.Discard) // caller prints the error; don't print it twice

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
