package cli

import (
	"fmt"

)

func Execute(args []string) error {

	if len(args) == 0 {
		Usage()
		return nil
	}

	switch args[0] {

	case "doctor":

		Doctor()
		return nil

	case "init":

		stack, err := ParseStack(args[1:])
		if err != nil {
			return err
		}

		return Init(stack)

	case "plan":

		stack, err := ParseStack(args[1:])
		if err != nil {
			return err
		}

		return Plan(stack)

	case "apply":

		stack, err := ParseStack(args[1:])
		if err != nil {
			return err
		}

		return Apply(stack)

	case "destroy":

		stack, err := ParseStack(args[1:])
		if err != nil {
			return err
		}

		return Destroy(stack)

	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}