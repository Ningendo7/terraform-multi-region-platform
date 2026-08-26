# Proves the cross-stack data pattern: bootstrap keeps local state (it
# creates the S3 backend, so it can't depend on that backend existing),
# so anything reading its outputs (e.g. the KMS key for encrypting flow
# logs, once modules/vpc is built) points at its state file directly.
data "terraform_remote_state" "bootstrap" {
  backend = "local"

  config = {
    path = "${path.module}/../../../bootstrap/terraform.tfstate"
  }
}

module "vpc" {
  source = "../../../../../modules/vpc"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  vpc_cidr = "10.0.0.0/16"
  az_count = 3

  # null until bootstrap is actually applied (its state has no outputs
  # yet) — flow logs still get created and encrypted with CloudWatch's
  # default key until then; the next plan after bootstrap is applied
  # picks up the real KMS key automatically, no code change needed.
  flow_log_kms_key_arn = try(data.terraform_remote_state.bootstrap.outputs.kms_key_arn, null)
}
