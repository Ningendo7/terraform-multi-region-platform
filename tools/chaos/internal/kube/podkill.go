package kube

import (
	"fmt"
	"math/rand"
	"strings"
)

// PodKill deletes one randomly chosen pod matching label within namespace,
// on the cluster identified by context. With dryRun set, it selects and
// prints the target without deleting anything.
func PodKill(context, namespace, label string, dryRun bool) error {

	out, err := Run(context, "get", "pods", "-n", namespace, "-l", label, "-o", "name")
	if err != nil {
		return err
	}

	pods := strings.Fields(out)
	if len(pods) == 0 {
		return fmt.Errorf("no pods matched namespace=%s label=%s", namespace, label)
	}

	target := pods[rand.Intn(len(pods))]

	if dryRun {
		fmt.Printf("[dry-run] would delete %s in namespace %s (matched %d candidates)\n", target, namespace, len(pods))
		return nil
	}

	fmt.Printf("deleting %s in namespace %s (matched %d candidates)\n", target, namespace, len(pods))

	if _, err := Run(context, "delete", target, "-n", namespace); err != nil {
		return err
	}

	fmt.Printf("deleted %s\n", target)
	return nil
}