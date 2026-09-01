package health

import (
	"context"
	"fmt"
	"time"
)

// Monitor checks every url. With interval <= 0 it checks once and
// returns an error if anything's unhealthy — what makes `probe check`
// usable as a CI gate. With interval > 0 it polls continuously,
// printing each round, until ctx is done (Ctrl-C, or --duration
// elapsing) — a long-running monitor never treats "still checking" as
// a failure.
func Monitor(ctx context.Context, urls []string, timeout, interval time.Duration) error {
	if interval <= 0 {
		return checkOnce(ctx, urls, timeout)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		_ = checkOnce(ctx, urls, timeout)

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
	}
}

func checkOnce(ctx context.Context, urls []string, timeout time.Duration) error {
	unhealthyCount := 0

	for _, url := range urls {
		result := Check(ctx, url, timeout)

		status := "healthy"
		if !result.Healthy {
			status = "unhealthy"
			unhealthyCount++
		}

		fmt.Printf("status=%s url=%s http_status=%d latency=%s err=%v\n",
			status, result.URL, result.StatusCode, result.Latency, result.Err)
	}

	if unhealthyCount > 0 {
		return fmt.Errorf("%d of %d targets unhealthy", unhealthyCount, len(urls))
	}

	return nil
}
