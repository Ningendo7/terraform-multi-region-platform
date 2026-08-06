package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// findRepoRoot walks up from the current working directory looking for
// the repo's terraform/ directory, so platformctl resolves stack paths
// the same way regardless of where it's invoked from.
func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if info, err := os.Stat(filepath.Join(dir, "terraform")); err == nil && info.IsDir() {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find repo root (no terraform/ directory found in any parent)")
		}
		dir = parent
	}
}

func ResolveStack(stack Stack) (string, error) {

	root, err := findRepoRoot()
	if err != nil {
		return "", err
	}

	switch stack.Scope {

	case ScopeBootstrap:
		return filepath.Join(
			root,
			"terraform",
			"bootstrap",
		), nil

	case ScopeGlobal:

		if stack.Name == "" {
			return "", fmt.Errorf("global stack requires a name")
		}

		return filepath.Join(
			root,
			"terraform",
			"global",
			stack.Name,
		), nil

	case ScopeRegion:

		if stack.Region == "" {
			return "", fmt.Errorf("region stack requires a region")
		}

		return filepath.Join(
			root,
			"terraform",
			"regions",
			stack.Region,
			stack.Name,
		), nil

	default:
		return "", fmt.Errorf("unknown stack scope: %s", stack.Scope)
	}
}
