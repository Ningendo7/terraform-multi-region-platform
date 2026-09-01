package safety

import "testing"

func TestRequireConfirmation(t *testing.T) {
	cases := []struct {
		name    string
		dryRun  bool
		yes     bool
		wantErr bool
	}{
		{name: "dry-run only", dryRun: true, yes: false},
		{name: "yes only", dryRun: false, yes: true},
		{name: "both set", dryRun: true, yes: true},
		{name: "neither set", dryRun: false, yes: false, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := RequireConfirmation(c.dryRun, c.yes)

			if c.wantErr && err == nil {
				t.Fatalf("RequireConfirmation(dryRun=%v, yes=%v) = nil, want error", c.dryRun, c.yes)
			}
			if !c.wantErr && err != nil {
				t.Fatalf("RequireConfirmation(dryRun=%v, yes=%v) unexpected error: %v", c.dryRun, c.yes, err)
			}
		})
	}
}
