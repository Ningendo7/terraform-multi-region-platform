package probecli

import (
	"flag"
	"fmt"
	"io"
	"time"
)

type CheckOptions struct {
	URLs     []string
	Timeout  time.Duration
	Interval time.Duration
	Duration time.Duration
}

func ParseCheckFlags(args []string) (CheckOptions, error) {

	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var urls stringSliceFlag
	fs.Var(&urls, "url", "target URL to check (repeatable — pass once per region)")

	var opts CheckOptions
	fs.DurationVar(&opts.Timeout, "timeout", 5*time.Second, "per-request timeout")
	fs.DurationVar(&opts.Interval, "interval", 0, "poll on this interval instead of checking once (0 = single check)")
	fs.DurationVar(&opts.Duration, "duration", 0, "stop continuous monitoring after this long (0 = run until Ctrl-C, only relevant with --interval)")

	if err := fs.Parse(args); err != nil {
		return CheckOptions{}, err
	}

	if len(urls) == 0 {
		return CheckOptions{}, fmt.Errorf("check requires at least one --url")
	}

	opts.URLs = urls
	return opts, nil
}

type RecoverOptions struct {
	URL      string
	Interval time.Duration
	Timeout  time.Duration
	MaxWait  time.Duration
	Settle   int
}

func ParseRecoverFlags(args []string) (RecoverOptions, error) {

	fs := flag.NewFlagSet("recover", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var opts RecoverOptions
	fs.StringVar(&opts.URL, "url", "", "target URL to watch (required)")
	fs.DurationVar(&opts.Interval, "interval", 2*time.Second, "how often to poll")
	fs.DurationVar(&opts.Timeout, "timeout", 5*time.Second, "per-request timeout")
	fs.DurationVar(&opts.MaxWait, "max-wait", 10*time.Minute, "give up if the full outage+recovery cycle doesn't finish within this")
	fs.IntVar(&opts.Settle, "settle", 3, "consecutive successful checks required before declaring recovery")

	if err := fs.Parse(args); err != nil {
		return RecoverOptions{}, err
	}

	if opts.URL == "" {
		return RecoverOptions{}, fmt.Errorf("recover requires --url")
	}

	return opts, nil
}

// stringSliceFlag lets --url be passed multiple times to build a list —
// the standard flag package has no built-in repeatable-flag type.
type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return fmt.Sprintf("%v", []string(*s))
}

func (s *stringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}
