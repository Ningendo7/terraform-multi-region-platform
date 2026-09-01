package awscli

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Run executes the aws CLI and returns its stdout. Mirrors internal/kube's
// Run for the same reason: shell out to the CLI the operator already has
// configured (real credentials, real profile) instead of pulling in the
// AWS SDK for a handful of read-only lookups.
func Run(ctx context.Context, region string, args ...string) (string, error) {
	fullArgs := append([]string{"--region", region}, args...)

	cmd := exec.CommandContext(ctx, "aws", fullArgs...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("aws %v failed: %w: %s", fullArgs, err, strings.TrimSpace(stderr.String()))
	}

	return stdout.String(), nil
}
