// tools/platformctl/internal/config/constants.go
package config

import "os"

const (
	// Terraform scopes
	ScopeBootstrap = "bootstrap"
	ScopeGlobal    = "global"
	ScopeRegion    = "region"
)

// ProjectName, Environment, and StateRegion identify the shared state
// backend this repo's own terraform/bootstrap provisions. They're read
// from the environment (defaulting to this repo's real values) rather
// than compiled in, so forking this repo to point at a different AWS
// account/project doesn't require editing and recompiling the CLI.
func ProjectName() string {
	return envOrDefault("PLATFORMCTL_PROJECT_NAME", "terraform-multi-region-platform")
}

func Environment() string {
	return envOrDefault("PLATFORMCTL_ENVIRONMENT", "prod")
}

func StateRegion() string {
	return envOrDefault("PLATFORMCTL_STATE_REGION", "us-east-1")
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
