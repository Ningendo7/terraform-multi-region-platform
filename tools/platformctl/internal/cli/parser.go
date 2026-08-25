package cli

import (
	"fmt"

	"github.com/Ningendo7/terraform-multi-region-platform/tools/platformctl/internal/config"
)

func ParseStack(args []string) (config.Stack, error) {

	if len(args) < 1 {
		return config.Stack{}, fmt.Errorf("missing stack information")
	}

	stack := config.Stack{
		Scope: args[0],
	}

	switch args[0] {

	case config.ScopeBootstrap:
		return stack, nil

	case config.ScopeGlobal:
		if len(args) < 2 {
			return config.Stack{}, fmt.Errorf("global requires stack name")
		}

		stack.Name = args[1]
		return stack, nil

	case config.ScopeRegion:

		if len(args) < 3 {
			return config.Stack{}, fmt.Errorf("region requires: <region> <stack>")
		}

		stack.Region = args[1]
		stack.Name = args[2]

		return stack, nil

	default:
		return config.Stack{}, fmt.Errorf("unknown scope: %s", args[0])
	}
}
