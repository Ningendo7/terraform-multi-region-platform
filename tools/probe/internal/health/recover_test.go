package health

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitForRecovery(t *testing.T) {
	var healthy atomic.Bool
	healthy.Store(true)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if healthy.Load() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}))
	defer server.Close()

	// simulate an outage starting shortly after the baseline check, and
	// recovering a bit after that
	go func() {
		time.Sleep(30 * time.Millisecond)
		healthy.Store(false)
		time.Sleep(60 * time.Millisecond)
		healthy.Store(true)
	}()

	result, err := WaitForRecovery(context.Background(), server.URL, 10*time.Millisecond, time.Second, time.Second, 2)
	if err != nil {
		t.Fatalf("WaitForRecovery() unexpected error: %v", err)
	}

	if !result.OutageDetected {
		t.Fatal("WaitForRecovery() did not detect the simulated outage")
	}

	if result.Downtime <= 0 {
		t.Fatalf("WaitForRecovery() downtime = %v, want > 0", result.Downtime)
	}
}

func TestWaitForRecoveryFailsIfUnhealthyAtStart(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	_, err := WaitForRecovery(context.Background(), server.URL, 10*time.Millisecond, time.Second, time.Second, 2)
	if err == nil {
		t.Fatal("WaitForRecovery() = nil error, want error when target is unhealthy at start")
	}
}
