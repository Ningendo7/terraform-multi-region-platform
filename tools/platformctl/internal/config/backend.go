package config

import "fmt"

func stateBucket() string {
	return fmt.Sprintf("%s-%s-terraform-state", ProjectName, Environment)
}

func lockTable() string {
	return fmt.Sprintf("%s-%s-terraform-lock", ProjectName, Environment)
}

// BackendConfig computes the terraform init -backend-config values for a
// stack. Bootstrap returns (nil, nil): it keeps local state, since it's
// what provisions the shared state bucket and lock table every other
// stack's backend config points at here.
func BackendConfig(stack Stack) (map[string]string, error) {

	if stack.Scope == ScopeBootstrap {
		return nil, nil
	}

	var key string

	switch stack.Scope {

	case ScopeGlobal:
		if stack.Name == "" {
			return nil, fmt.Errorf("global stack requires a name")
		}
		key = fmt.Sprintf("global/%s/terraform.tfstate", stack.Name)

	case ScopeRegion:
		if stack.Region == "" {
			return nil, fmt.Errorf("region stack requires a region")
		}
		key = fmt.Sprintf("regions/%s/%s/terraform.tfstate", stack.Region, stack.Name)

	default:
		return nil, fmt.Errorf("unknown stack scope: %s", stack.Scope)
	}

	return map[string]string{
		"bucket":         stateBucket(),
		"key":            key,
		"region":         StateRegion,
		"dynamodb_table": lockTable(),
		"encrypt":        "true",
	}, nil
}
