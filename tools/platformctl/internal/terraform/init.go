package terraform

import (
	"fmt"
	"sort"
)

func Init(directory string, backendConfig map[string]string) error {

	args := []string{"init"}

	keys := make([]string, 0, len(backendConfig))
	for k := range backendConfig {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		args = append(args, fmt.Sprintf("-backend-config=%s=%s", k, backendConfig[k]))
	}

	return Execute(directory, args...)
}
