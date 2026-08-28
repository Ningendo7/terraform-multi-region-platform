package chaoscli

import (
	"flag"
	"fmt"
)

type PodKillOptions struct {
	Context   string
	Namespace string
	Label     string
	DryRun    bool
}

func ParsePodKillFlags(args []string) (PodKillOptions, error) {

	fs := flag.NewFlagSet("pod-kill", flag.ContinueOnError)

	var opts PodKillOptions
	fs.StringVar(&opts.Context, "context", "", "kubeconfig context of the target cluster (required)")
	fs.StringVar(&opts.Namespace, "namespace", "", "namespace to target (required)")
	fs.StringVar(&opts.Label, "label", "", "label selector for candidate pods, e.g. app=demo-app (required)")
	fs.BoolVar(&opts.DryRun, "dry-run", false, "select a target and print it without deleting anything")

	if err := fs.Parse(args); err != nil {
		return PodKillOptions{}, err
	}

	if opts.Context == "" || opts.Namespace == "" || opts.Label == "" {
		return PodKillOptions{}, fmt.Errorf("pod-kill requires --context, --namespace, and --label")
	}

	return opts, nil
}