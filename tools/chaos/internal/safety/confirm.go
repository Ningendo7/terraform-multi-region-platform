package safety

import "fmt"

// RequireConfirmation is the shared gate every destructive experiment
// runs through before touching anything: refuse to proceed unless the
// caller passed --dry-run (nothing will actually happen) or --yes
// (explicit, deliberate opt-in to the real thing). No interactive
// prompt — chaos experiments are often run from automation, and a
// prompt that blocks on stdin is worse there than requiring an
// explicit flag.
func RequireConfirmation(dryRun bool, yes bool) error {
	if dryRun || yes {
		return nil
	}
	return fmt.Errorf("refusing to run without --dry-run or --yes")
}
