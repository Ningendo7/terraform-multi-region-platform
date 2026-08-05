output "kms_key_arn" {
  description = "The Bootstrap KMS key ARN."
  value       = module.kms.key_arn
}

output "kms_key_alias" {
  description = "The alias of the Bootstrap KMS key."
  value       = module.kms.alias
}

output "state_bucket" {
  description = "The name of the Terraform state S3 bucket."
  value       = module.state_bucket.bucket_name
}

output "terraform_lock_table" {
  description = "The name of the Terraform lock DynamoDB table."
  value       = module.terraform_lock.table_name
}