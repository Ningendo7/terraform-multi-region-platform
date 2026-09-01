package probecli

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/probe/internal/health"
)

func Execute(args []string) error {

	if len(args) == 0 {
		Usage()
		return nil
	}

	switch args[0] {

	case "check":
		return runCheck(args[1:])

	case "recover":
		return runRecover(args[1:])

	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runCheck(args []string) error {

	opts, err := ParseCheckFlags(args)
	if err != nil {
		return err
	}

	// Ctrl-C always stops cleanly, even in continuous mode.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	switch {
	case opts.Interval <= 0:
		// single-shot: bound execution to just the checks themselves
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.Timeout+10*time.Second)
		defer cancel()
	case opts.Duration > 0:
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.Duration)
		defer cancel()
		// else: continuous and unbounded — runs until Ctrl-C
	}

	return health.Monitor(ctx, opts.URLs, opts.Timeout, opts.Interval)
}

func runRecover(args []string) error {

	opts, err := ParseRecoverFlags(args)
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, opts.MaxWait+opts.Timeout)
	defer cancel()

	result, err := health.WaitForRecovery(ctx, opts.URL, opts.Interval, opts.Timeout, opts.MaxWait, opts.Settle)
	if err != nil {
		return err
	}

	fmt.Printf("outage_start=%s recovered_at=%s downtime=%s\n",
		result.OutageStart.Format("15:04:05"), result.RecoveredAt.Format("15:04:05"), result.Downtime)

	return nil
}
