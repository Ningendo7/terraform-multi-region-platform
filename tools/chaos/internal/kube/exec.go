package kube

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Run executes kubectl against the given context and returns its stdout.
// Stderr is passed through directly, matching how internal/terraform
// wraps the terraform CLI.
func Run(context string, args ...string) (string, error) {
	fullArgs := append([]string{"--context", context}, args...)

	cmd := exec.Command("kubectl", fullArgs...)
	cmd.Stderr = os.Stderr

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("kubectl %v failed: %w", fullArgs, err)
	}

	return stdout.String(), nil
}