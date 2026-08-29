// tools/platformctl/internal/terraform/init_test.go
package terraform

import (
	"slices"
	"testing"
)

func TestBuildInitArgs(t *testing.T) {
	cases := []struct {
		name          string
		backendConfig map[string]string
		want          []string
	}{
		{
			name: "nil backend config",
			want: []string{"init"},
		},
		{
			name:          "empty backend config",
			backendConfig: map[string]string{},
			want:          []string{"init"},
		},
		{
			name:          "single key",
			backendConfig: map[string]string{"bucket": "my-bucket"},
			want:          []string{"init", "-backend-config=bucket=my-bucket"},
		},
		{
			name: "multiple keys are sorted deterministically",
			backendConfig: map[string]string{
				"region": "us-east-1",
				"bucket": "my-bucket",
				"key":    "path/to/state",
			},
			want: []string{
				"init",
				"-backend-config=bucket=my-bucket",
				"-backend-config=key=path/to/state",
				"-backend-config=region=us-east-1",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := buildInitArgs(c.backendConfig)

			if !slices.Equal(got, c.want) {
				t.Fatalf("buildInitArgs(%+v) = %v, want %v", c.backendConfig, got, c.want)
			}
		})
	}
}
