package cli

import (
	"fmt"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/checks"
)

func Doctor() {

	fmt.Println("Platform Environment Check")
	fmt.Println("----------------------------")

	tools := []string{
		"terraform", 
		"aws", 
		"kubectl", 
		"helm", 
		"git",
	}

	for _, tool := range tools {
		checks.CheckCommand(tool)
	}
}