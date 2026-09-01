package terraform

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Execute(ctx context.Context, directory string, arguments ...string) error {
	cmd := exec.CommandContext(ctx, "terraform", arguments...)
	cmd.Dir = directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// On timeout/cancellation, ask terraform to shut down cleanly (SIGINT,
	// same as a human hitting Ctrl-C) instead of killing it outright —
	// terraform catches SIGINT and tries to finish its in-flight API call
	// and write state before exiting, so state stays consistent with
	// what's actually in AWS. Only force-kill if it doesn't respond
	// within WaitDelay.
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}
	cmd.WaitDelay = 15 * time.Second

	if err := cmd.Run(); err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return fmt.Errorf("terraform %v timed out and was asked to shut down cleanly — check `terraform state list` in %s before retrying: %w", arguments, directory, err)
		}
		return fmt.Errorf("terraform %v failed: %w", arguments, err)
	}

	return nil
}
