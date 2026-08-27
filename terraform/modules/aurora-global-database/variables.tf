variable "project_name" {
  description = "The name of the project."
  type        = string
}

variable "environment" {
  description = "The environment name."
  type        = string
}

variable "identifier" {
  description = "Logical database name — gets prefixed with project/environment, matching every other resource in this repo."
  type        = string
}

variable "is_primary" {
  description = "Whether this call creates the Aurora Global Database (true) or attaches a secondary regional cluster to an existing one (false)."
  type        = bool
  default     = true
}

variable "global_cluster_identifier" {
  description = "ID of the existing global cluster to attach to. Required when is_primary is false; ignored when true (the module creates its own)."
  type        = string
  default     = null
}

variable "engine_version" {
  description = "Aurora PostgreSQL engine version."
  type        = string
  default     = "16.4"
}

variable "database_name" {
  description = "Initial database name. Only meaningful when is_primary is true — secondary regions replicate it, they don't set their own."
  type        = string
  default     = null
}

variable "master_username" {
  description = "Master username. Only meaningful when is_primary is true."
  type        = string
  default     = "app_admin"
}

variable "instance_class" {
  description = "Instance class for every instance in this region's cluster."
  type        = string
  default     = "db.r6g.large"
}

variable "instance_count" {
  description = "Number of instances (writer + readers) in this region's cluster."
  type        = number
  default     = 1
}

variable "vpc_id" {
  description = "VPC ID this region's cluster deploys into."
  type        = string
}

variable "subnet_ids" {
  description = "Subnet IDs for the DB subnet group — private subnets."
  type        = list(string)
}

variable "allowed_security_group_ids" {
  description = "Security group IDs allowed to reach the database on the Postgres port (e.g. an EKS cluster's security group)."
  type        = list(string)
  default     = []
}

variable "kms_key_arn" {
  description = "KMS key ARN for storage encryption. If null, RDS's own AWS-managed key is used instead."
  type        = string
  default     = null
}

variable "backup_retention_period" {
  description = "Backup retention period, in days."
  type        = number
  default     = 7
}

variable "deletion_protection" {
  description = "Whether to enable deletion protection — a live AWS setting; must be explicitly disabled before the cluster can be destroyed."
  type        = bool
  default     = true
}

variable "skip_final_snapshot" {
  description = "Whether to skip taking a final snapshot on destroy. Default false — take the snapshot."
  type        = bool
  default     = false
}

variable "iam_database_authentication_enabled" {
  description = "Whether to allow IAM-based authentication to the database, alongside the master password — lets IRSA-derived credentials connect without a static password."
  type        = bool
  default     = true
}