package health

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	healthy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer healthy.Close()

	unhealthy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer unhealthy.Close()

	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer slow.Close()

	cases := []struct {
		name        string
		url         string
		timeout     time.Duration
		wantHealthy bool
	}{
		{name: "2xx is healthy", url: healthy.URL, timeout: time.Second, wantHealthy: true},
		{name: "5xx is unhealthy", url: unhealthy.URL, timeout: time.Second, wantHealthy: false},
		{name: "timeout is unhealthy", url: slow.URL, timeout: 10 * time.Millisecond, wantHealthy: false},
		{name: "unreachable is unhealthy", url: "http://127.0.0.1:1", timeout: time.Second, wantHealthy: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Check(context.Background(), c.url, c.timeout)

			if result.Healthy != c.wantHealthy {
				t.Fatalf("Check(%s) healthy = %v, want %v (err: %v)", c.url, result.Healthy, c.wantHealthy, result.Err)
			}
		})
	}
}
