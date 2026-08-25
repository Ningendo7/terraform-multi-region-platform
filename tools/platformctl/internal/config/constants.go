package config

const (

	// Terraform scopes

	ScopeBootstrap = "bootstrap"
	ScopeGlobal    = "global"
	ScopeRegion    = "region"

	// Environment is fixed to "prod" for now — only prod is scaffolded.
	// Deliberately not a CLI flag yet: see plan notes on deferring --env
	// until a second environment actually exists.
	Environment = "prod"

	// ProjectName and StateRegion identify the shared state backend that
	// terraform/environments/prod/bootstrap provisions. StateRegion is
	// where the state bucket itself lives — independent of which AWS
	// region any given stack's resources are deployed into.
	ProjectName = "terraform-multi-region-platform"
	StateRegion = "us-east-1"
)
