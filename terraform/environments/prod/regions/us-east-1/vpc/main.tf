module "flow_log_kms" {
  source = "../../../../../modules/kms"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  description            = "VPC flow logs KMS key — us-east-1"
  name_suffix            = "vpc-flow-logs"
  enable_cloudwatch_logs = true
}

module "vpc" {
  source = "../../../../../modules/vpc"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  vpc_cidr = "10.0.0.0/16"
  az_count = 3

  flow_log_kms_key_arn = module.flow_log_kms.key_arn
}
