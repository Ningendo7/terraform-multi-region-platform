// tools/platformctl/internal/cli/doctor.go
package cli

import (
	"fmt"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/checks"
)

func Doctor() error {

	fmt.Println("Platform Environment Check")
	fmt.Println("----------------------------")

	tools := []string{"terraform", "aws", "kubectl", "helm", "git"}

	allFound := true
	for _, tool := range tools {
		if !checks.CheckCommand(tool) {
			allFound = false
		}
	}

	if !allFound {
		return fmt.Errorf("one or more required tools are missing")
	}

	return nil
}
