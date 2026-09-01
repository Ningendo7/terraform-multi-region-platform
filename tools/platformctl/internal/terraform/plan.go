package terraform

import (
	"context"
	"time"
)

const PlanTimeout = 10 * time.Minute

func Plan(directory string) error {
	ctx, cancel := context.WithTimeout(context.Background(), PlanTimeout)
	defer cancel()

	return Execute(ctx, directory, "plan", "-lock-timeout=30s")
}
