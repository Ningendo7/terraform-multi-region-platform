// tools/platformctl/internal/checks/tools.go
package checks

import (
	"fmt"
	"os/exec"
)

// CheckCommand reports whether command is on PATH, printing a line either
// way, and returns whether it was found — so callers that need `doctor`
// to work as a CI gate (not just a human-readable report) can act on it.
func CheckCommand(command string) bool {
	_, err := exec.LookPath(command)

	if err != nil {
		fmt.Printf("x %s not found\n", command)
		return false
	}

	fmt.Printf(" %s installed\n", command)
	return true
}
