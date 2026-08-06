package terraform

import (
	"fmt"
	"os/exec"
)

func Execute(directory string, arguments ...string) error {
	cmd := exec.Command("terraform", arguments...)
	cmd.Dir = directory

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"terraform %v failed:\n%s\n%w", 
			arguments, 
			string(output), 
			err,
		)
	}
/*
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute terraform command: %w", arguments, err)
	}
*/
	return nil
}