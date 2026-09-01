package chaoscli

import "testing"

func TestParsePodKillFlags(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		want    PodKillOptions
		wantErr bool
	}{
		{
			name: "dry-run",
			args: []string{"--context", "us-east-1", "--namespace", "demo", "--label", "app=demo-app", "--dry-run"},
			want: PodKillOptions{Context: "us-east-1", Namespace: "demo", Label: "app=demo-app", DryRun: true},
		},
		{
			name: "real delete with --yes",
			args: []string{"--context", "us-east-1", "--namespace", "demo", "--label", "app=demo-app", "--yes"},
			want: PodKillOptions{Context: "us-east-1", Namespace: "demo", Label: "app=demo-app", Yes: true},
		},
		{
			name: "explicit seed",
			args: []string{"--context", "us-east-1", "--namespace", "demo", "--label", "app=demo-app", "--dry-run", "--seed", "42"},
			want: PodKillOptions{Context: "us-east-1", Namespace: "demo", Label: "app=demo-app", DryRun: true, Seed: 42},
		},
		{
			name:    "missing context",
			args:    []string{"--namespace", "demo", "--label", "app=demo-app", "--dry-run"},
			wantErr: true,
		},
		{
			name:    "missing namespace",
			args:    []string{"--context", "us-east-1", "--label", "app=demo-app", "--dry-run"},
			wantErr: true,
		},
		{
			name:    "missing label",
			args:    []string{"--context", "us-east-1", "--namespace", "demo", "--dry-run"},
			wantErr: true,
		},
		{
			name:    "no args",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "unknown flag",
			args:    []string{"--bogus", "value"},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParsePodKillFlags(c.args)

			if c.wantErr {
				if err == nil {
					t.Fatalf("ParsePodKillFlags(%v) = %+v, want error", c.args, got)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParsePodKillFlags(%v) unexpected error: %v", c.args, err)
			}

			if got != c.want {
				t.Fatalf("ParsePodKillFlags(%v) = %+v, want %+v", c.args, got, c.want)
			}
		})
	}
}
