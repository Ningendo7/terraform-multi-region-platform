package health

import (
	"context"
	"fmt"
	"time"
)

// RecoveryResult reports what happened during a waitForRecovery run.
type RecoveryResult struct {
	OutageDetected bool
	OutageStart    time.Time
	RecoveredAt    time.Time
	Downtime       time.Duration
}

// WaitForRecovery polls url every interval, starting from a required
// healthy baseline, and measures how long it takes to go unhealthy and
// come back — the actual number behind a resilience claim ("failover
// took 47 seconds") instead of a guess.
//
// settle consecutive successes are required before declaring recovery,
// so one flaky response mid-outage doesn't get reported as the moment
// things came back. Meant to be started before triggering a chaos
// experiment in another terminal; maxWait bounds the whole cycle.
func WaitForRecovery(ctx context.Context, url string, interval, timeout, maxWait time.Duration, settle int) (RecoveryResult, error) {
	baseline := Check(ctx, url, timeout)
	if !baseline.Healthy {
		return RecoveryResult{}, fmt.Errorf("%s is not healthy at the start — nothing to measure a recovery from (status=%d err=%v)",
			url, baseline.StatusCode, baseline.Err)
	}

	fmt.Printf("baseline healthy, watching %s for an outage (up to %s)...\n", url, maxWait)

	deadline := time.Now().Add(maxWait)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var result RecoveryResult
	consecutiveSuccesses := 0

	for {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		case <-ticker.C:
		}

		if time.Now().After(deadline) {
			if !result.OutageDetected {
				return result, fmt.Errorf("no outage detected within %s", maxWait)
			}
			return result, fmt.Errorf("outage detected at %s but hadn't recovered within %s", result.OutageStart.Format(time.RFC3339), maxWait)
		}

		check := Check(ctx, url, timeout)

		if !check.Healthy {
			if !result.OutageDetected {
				result.OutageDetected = true
				result.OutageStart = time.Now()
				fmt.Printf("action=outage-detected url=%s at=%s\n", url, result.OutageStart.Format(time.RFC3339))
			}
			consecutiveSuccesses = 0
			continue
		}

		if !result.OutageDetected {
			continue
		}

		consecutiveSuccesses++
		if consecutiveSuccesses >= settle {
			result.RecoveredAt = time.Now()
			result.Downtime = result.RecoveredAt.Sub(result.OutageStart)
			fmt.Printf("action=recovered url=%s at=%s downtime=%s\n",
				url, result.RecoveredAt.Format(time.RFC3339), result.Downtime)
			return result, nil
		}
	}
}
