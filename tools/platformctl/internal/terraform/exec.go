package terraform

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(directory string, arguments ...string) error {
	cmd := exec.Command("terraform", arguments...)
	cmd.Dir = directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("terraform %v failed: %w", arguments, err)
	}

	return nil
}
