package kube

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Run executes kubectl against the given context and returns its stdout.
// Stderr is captured into the returned error (not just streamed) so callers
// — and anything logging this — see the real reason a call failed. Callers
// control cancellation/timeout via ctx; this package never hangs silently.
func Run(ctx context.Context, kubecontext string, args ...string) (string, error) {
	fullArgs := append([]string{"--context", kubecontext}, args...)

	cmd := exec.CommandContext(ctx, "kubectl", fullArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("kubectl %v failed: %w: %s", fullArgs, err, strings.TrimSpace(stderr.String()))
	}

	return stdout.String(), nil
}
