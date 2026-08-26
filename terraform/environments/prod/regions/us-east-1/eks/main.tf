# vpc uses the real S3 backend (unlike bootstrap's local state), so this
# reads the same bucket/key platformctl computed for it.
data "terraform_remote_state" "vpc" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-east-1/vpc/terraform.tfstate"
    region = "us-east-1"
  }
}

module "eks" {
  source = "../../../../../modules/eks"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  # No try()/fallback here on purpose — an EKS cluster genuinely cannot
  # exist without a VPC, so this plan is *supposed* to fail until the
  # vpc stack is actually applied. Papering over that with a default
  # would just produce a plan that isn't meaningful.
  vpc_id             = data.terraform_remote_state.vpc.outputs.vpc_id
  private_subnet_ids = data.terraform_remote_state.vpc.outputs.private_subnet_ids
  public_subnet_ids  = data.terraform_remote_state.vpc.outputs.public_subnet_ids
}
