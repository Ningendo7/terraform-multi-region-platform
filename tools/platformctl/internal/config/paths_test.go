package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolveStack(t *testing.T) {
	cases := []struct {
		name       string
		stack      Stack
		wantSuffix string
		wantErr    bool
	}{
		{
			name:       "bootstrap",
			stack:      Stack{Scope: ScopeBootstrap},
			wantSuffix: filepath.Join("terraform", "environments", "prod", "bootstrap"),
		},
		{
			name:       "global",
			stack:      Stack{Scope: ScopeGlobal, Name: "iam"},
			wantSuffix: filepath.Join("terraform", "environments", "prod", "global", "iam"),
		},
		{
			name:    "global missing name",
			stack:   Stack{Scope: ScopeGlobal},
			wantErr: true,
		},
		{
			name:       "region",
			stack:      Stack{Scope: ScopeRegion, Region: "us-east-1", Name: "eks"},
			wantSuffix: filepath.Join("terraform", "environments", "prod", "regions", "us-east-1", "eks"),
		},
		{
			name:    "region missing region",
			stack:   Stack{Scope: ScopeRegion, Name: "eks"},
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
			got, err := ResolveStack(c.stack)

			if c.wantErr {
				if err == nil {
					t.Fatalf("ResolveStack(%+v) = %q, want error", c.stack, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("ResolveStack(%+v) unexpected error: %v", c.stack, err)
			}

			if !strings.HasSuffix(got, c.wantSuffix) {
				t.Fatalf("ResolveStack(%+v) = %q, want suffix %q", c.stack, got, c.wantSuffix)
			}

			if !filepath.IsAbs(got) {
				t.Fatalf("ResolveStack(%+v) = %q, want an absolute path", c.stack, got)
			}
		})
	}
}

func TestFindRepoRoot(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("findRepoRoot() unexpected error: %v", err)
	}

	if info, err := os.Stat(filepath.Join(root, "terraform")); err != nil || !info.IsDir() {
		t.Fatalf("findRepoRoot() = %q, does not contain a terraform/ directory", root)
	}
}
