module "flow_log_kms" {
  source = "../../../../../modules/kms"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  description            = "VPC flow logs KMS key — us-east-2"
  name_suffix            = "vpc-flow-logs"
  enable_cloudwatch_logs = true
}

module "vpc" {
  source = "../../../../../modules/vpc"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  # Distinct from us-east-1's 10.0.0.0/16 — no functional requirement
  # to avoid overlap since these VPCs are never peered (see the
  # multi-region networking discussion — independent, active-active
  # regions don't talk to each other), but non-overlapping ranges cost
  # nothing and keep the door open if that ever changes.
  vpc_cidr = "10.1.0.0/16"
  az_count = 3

  flow_log_kms_key_arn = module.flow_log_kms.key_arn

  # IAM is account-wide, not regional — without this, module.vpc's
  # flow_logs IAM role would collide with us-east-1's.
  name_suffix = "us-east-2"

  public_subnet_tags = {
    "kubernetes.io/cluster/terraform-multi-region-platform-prod" = "shared"
    "kubernetes.io/role/elb"                                     = "1"
  }

  private_subnet_tags = {
    "kubernetes.io/cluster/terraform-multi-region-platform-prod" = "shared"
    "kubernetes.io/role/internal-elb"                            = "1"
  }
}
