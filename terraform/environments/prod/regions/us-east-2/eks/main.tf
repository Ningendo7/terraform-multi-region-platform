# The state bucket itself always lives in us-east-1 regardless of
# which region's stack this is — bucket/region below is "where state
# is stored", unrelated to which AWS region this stack's resources
# deploy into. Only the key changes.
data "terraform_remote_state" "vpc" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-east-2/vpc/terraform.tfstate"
    region = "us-east-1"
  }
}

module "eks" {
  source = "../../../../../modules/eks"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  vpc_id             = data.terraform_remote_state.vpc.outputs.vpc_id
  private_subnet_ids = data.terraform_remote_state.vpc.outputs.private_subnet_ids
  public_subnet_ids  = data.terraform_remote_state.vpc.outputs.public_subnet_ids

  # IAM is account-wide, not regional — without this, every IAM role
  # this module creates would collide with us-east-1's. Abbreviated
  # ("use2" not "us-east-2") because IAM role names cap at 64 chars —
  # the full region name pushed the longest one (karpenter-controller)
  # to 67.
  name_suffix = "use2"

  # Same shared zone as us-east-1 — one hosted zone for the whole
  # platform, not one per region.
  route53_zone_arn = "arn:aws:route53:::hostedzone/Z02533231WT3LO3AOOYF2"
}
