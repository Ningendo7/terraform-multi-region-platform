package checks

import (
	"fmt"
	"os/exec"
)

func CheckCommand(command string) {
	_, err := exec.LookPath(command)

	if err != nil {
		fmt.Printf("x %s not found\n", command)
		return
	}

	fmt.Printf(" %s installed\n", command)
}
