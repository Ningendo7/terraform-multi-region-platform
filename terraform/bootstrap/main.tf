module "kms" {
  source = "../modules/kms"

  project_name = var.project_name
  environment  = var.environment

  description = "Terraform bootstrap KMS Key"

}

module "state_bucket" {
  source = "../modules/s3"

  bucket_name = "${var.project_name}-terraform-state"
  kms_key_arn = module.kms.key_arn
  environment = var.environment
}