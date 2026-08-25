package cli

import (
	"fmt"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/config"
	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/terraform"
)

func Apply(stack config.Stack) error {

	path, err := config.ResolveStack(stack)
	if err != nil {
		return err
	}

	fmt.Printf("Applying stack: %s\n", path)

	return terraform.Apply(path)
}
