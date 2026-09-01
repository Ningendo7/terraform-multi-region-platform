package terraform

import (
	"context"
	"time"
)

const DestroyTimeout = 45 * time.Minute

func Destroy(directory string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DestroyTimeout)
	defer cancel()

	return Execute(ctx, directory, "destroy", "-lock-timeout=30s")
}
