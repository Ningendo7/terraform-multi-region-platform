package config

import "testing"

func TestBackendConfig(t *testing.T) {
	cases := []struct {
		name    string
		stack   Stack
		wantKey string
		wantErr bool
		wantNil bool
	}{
		{
			name:    "bootstrap has no backend config",
			stack:   Stack{Scope: ScopeBootstrap},
			wantNil: true,
		},
		{
			name:    "global",
			stack:   Stack{Scope: ScopeGlobal, Name: "iam"},
			wantKey: "global/iam/terraform.tfstate",
		},
		{
			name:    "global missing name",
			stack:   Stack{Scope: ScopeGlobal},
			wantErr: true,
		},
		{
			name:    "region",
			stack:   Stack{Scope: ScopeRegion, Region: "us-east-1", Name: "vpc"},
			wantKey: "regions/us-east-1/vpc/terraform.tfstate",
		},
		{
			name:    "region missing region",
			stack:   Stack{Scope: ScopeRegion, Name: "vpc"},
			wantErr: true,
		},
		{
			name:    "unknown scope",
			stack:   Stack{Scope: "nonsense"},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := BackendConfig(c.stack)

			if c.wantErr {
				if err == nil {
					t.Fatalf("BackendConfig(%+v) = %+v, want error", c.stack, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("BackendConfig(%+v) unexpected error: %v", c.stack, err)
			}

			if c.wantNil {
				if got != nil {
					t.Fatalf("BackendConfig(%+v) = %+v, want nil", c.stack, got)
				}
				return
			}

			if got["key"] != c.wantKey {
				t.Fatalf("BackendConfig(%+v) key = %q, want %q", c.stack, got["key"], c.wantKey)
			}

			for _, field := range []string{"bucket", "region", "use_lockfile", "encrypt"} {
				if got[field] == "" {
					t.Fatalf("BackendConfig(%+v) missing %q", c.stack, field)
				}
			}
		})
	}
}
