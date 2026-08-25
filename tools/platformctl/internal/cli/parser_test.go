package cli

import (
	"testing"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/config"
)

func TestParseStack(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		want    config.Stack
		wantErr bool
	}{
		{
			name: "bootstrap",
			args: []string{"bootstrap"},
			want: config.Stack{Scope: config.ScopeBootstrap},
		},
		{
			name: "global",
			args: []string{"global", "iam"},
			want: config.Stack{Scope: config.ScopeGlobal, Name: "iam"},
		},
		{
			name:    "global missing name",
			args:    []string{"global"},
			wantErr: true,
		},
		{
			name: "region",
			args: []string{"region", "us-east-1", "eks"},
			want: config.Stack{Scope: config.ScopeRegion, Region: "us-east-1", Name: "eks"},
		},
		{
			name:    "region missing stack",
			args:    []string{"region", "us-east-1"},
			wantErr: true,
		},
		{
			name:    "no args",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "unknown scope",
			args:    []string{"nonsense"},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseStack(c.args)

			if c.wantErr {
				if err == nil {
					t.Fatalf("ParseStack(%v) = %+v, want error", c.args, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseStack(%v) unexpected error: %v", c.args, err)
			}

			if got != c.want {
				t.Fatalf("ParseStack(%v) = %+v, want %+v", c.args, got, c.want)
			}
		})
	}
}
