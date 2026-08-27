variable "project_name" {
  description = "The name of the project."
  type        = string
}

variable "environment" {
  description = "The environment name."
  type        = string
}

variable "table_name" {
  description = "Logical table name — gets prefixed with project/environment, matching every other resource in this repo."
  type        = string
}

variable "billing_mode" {
  description = "DynamoDB billing mode."
  type        = string
  default     = "PAY_PER_REQUEST"
}

variable "hash_key" {
  description = "Partition key attribute name."
  type        = string
}

variable "range_key" {
  description = "Sort key attribute name, if any."
  type        = string
  default     = null
}

variable "attributes" {
  description = "Attribute definitions for every key/index attribute (DynamoDB only requires declaring attributes actually used as keys, not every field)."
  type = list(object({
    name = string
    type = string
  }))
}

variable "replica_regions" {
  description = "Additional AWS regions to replicate this table into, beyond wherever the module itself is applied. Empty by default — a table with zero replicas is still valid, and becomes a real Global Table the moment a region is added here."
  type        = list(string)
  default     = []
}

variable "kms_key_arn" {
  description = "KMS key ARN for encryption at rest. If null, DynamoDB's own AWS-managed key is used instead."
  type        = string
  default     = null
}

variable "point_in_time_recovery_enabled" {
  description = "Whether to enable point-in-time recovery."
  type        = bool
  default     = true
}

variable "deletion_protection_enabled" {
  description = "Whether to enable DynamoDB's own deletion protection (a live AWS setting, not just a Terraform lifecycle guard — must be explicitly disabled before the table can be destroyed)."
  type        = bool
  default     = true
}
