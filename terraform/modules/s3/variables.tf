variable "bucket_name" {
  description = "The name of the S3 bucket."
  type        = string

}

variable "kms_key_arn" {
  description = "The ARN of the KMS key to use for the S3 bucket."
  type        = string
}

variable "environment" {
  description = "The environment name."
  type        = string

}