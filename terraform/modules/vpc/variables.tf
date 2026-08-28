variable "project_name" {
  description = "The name of the project."
  type        = string
}

variable "environment" {
  description = "The environment name."
  type        = string
}

variable "vpc_cidr" {
  description = "The CIDR block for the VPC."
  type        = string
  default     = "10.0.0.0/16"
}

variable "az_count" {
  description = "Number of Availability Zones to spread subnets and NAT gateways across."
  type        = number
  default     = 3
}

variable "enable_flow_logs" {
  description = "Whether to enable VPC flow logs to CloudWatch Logs."
  type        = bool
  default     = true
}

variable "flow_log_kms_key_arn" {
  description = "KMS key ARN used to encrypt the flow log CloudWatch Logs group. If null, CloudWatch's default encryption is used instead."
  type        = string
  default     = null
}

variable "flow_log_retention_days" {
  description = "CloudWatch Logs retention period for VPC flow logs, in days."
  type        = number
  default     = 90
}

variable "public_subnet_tags" {
  description = "Additional tags applied only to public subnets (e.g. Kubernetes ELB discovery tags, added once EKS exists)."
  type        = map(string)
  default     = {}
}

variable "private_subnet_tags" {
  description = "Additional tags applied only to private subnets (e.g. Kubernetes internal-ELB discovery tags, added once EKS exists)."
  type        = map(string)
  default     = {}
}

variable "name_suffix" {
  description = "Optional suffix appended only to this module's IAM role name (flow_logs), which is account-wide unique unlike everything else here (VPC/subnet/NAT tags, CloudWatch log group names — all regional, no collision risk). Needed once this module is called more than once in the same account (e.g. a second region) with the same project_name/environment. Leave empty to preserve existing naming."
  type        = string
  default     = ""
}