// tools/platformctl/internal/terraform/apply.go
package terraform

import (
	"context"
	"time"
)

// 45 minutes covers real long-pole operations this platform has hit in
// practice (EKS cluster create/destroy can take 15-20 minutes) with
// comfortable headroom, while still guaranteeing the tool eventually
// gives up and reports a hung call instead of blocking forever. Safe to
// re-run afterward — a timed-out apply/destroy just resumes from
// whatever state was last written.
const ApplyTimeout = 45 * time.Minute

func Apply(directory string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ApplyTimeout)
	defer cancel()

	return Execute(ctx, directory, "apply", "-lock-timeout=30s")
}
