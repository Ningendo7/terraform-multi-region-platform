package cli

import (
	"fmt"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/config"
	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/terraform"
)

func Init(stack config.Stack) error {

	path, err := config.ResolveStack(stack)
	if err != nil {
		return err
	}

	fmt.Printf("Initializing stack: %s\n", path)

	return terraform.Init(path)
}
	