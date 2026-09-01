package health

import (
	"context"
	"net/http"
	"time"
)

// Result is the outcome of a single health check against one target.
type Result struct {
	URL        string
	Healthy    bool
	StatusCode int
	Latency    time.Duration
	Err        error
}

// Check performs one GET request against url and reports whether it
// succeeded within timeout. A response under 400 counts as healthy —
// anything else, including a network error or timeout, does not.
func Check(ctx context.Context, url string, timeout time.Duration) Result {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Result{URL: url, Err: err, Latency: time.Since(start)}
	}

	resp, err := http.DefaultClient.Do(req)
	latency := time.Since(start)

	if err != nil {
		return Result{URL: url, Err: err, Latency: latency}
	}
	defer resp.Body.Close()

	return Result{
		URL:        url,
		Healthy:    resp.StatusCode < 400,
		StatusCode: resp.StatusCode,
		Latency:    latency,
	}
}
