package terraform

import (
	"context"
	"fmt"
	"sort"
	"time"
)

const InitTimeout = 5 * time.Minute

// buildInitArgs turns a backend-config map into deterministically-ordered
// -backend-config flags — pulled out as its own pure function so it's
// testable without shelling out to a real terraform binary.
func buildInitArgs(backendConfig map[string]string) []string {

	args := []string{"init"}

	keys := make([]string, 0, len(backendConfig))
	for k := range backendConfig {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		args = append(args, fmt.Sprintf("-backend-config=%s=%s", k, backendConfig[k]))
	}

	return args
}

func Init(directory string, backendConfig map[string]string) error {

	ctx, cancel := context.WithTimeout(context.Background(), InitTimeout)
	defer cancel()

	return Execute(ctx, directory, buildInitArgs(backendConfig)...)
}
