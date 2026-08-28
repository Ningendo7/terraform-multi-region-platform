variable "project_name" {
  description = "The name of the project."
  type        = string
}

variable "environment" {
  description = "The environment name."
  type        = string
}

variable "kubernetes_version" {
  description = "EKS Kubernetes version."
  type        = string
  default     = "1.31"
}

variable "vpc_id" {
  description = "VPC ID the cluster and its nodes deploy into."
  type        = string
}

variable "private_subnet_ids" {
  description = "Private subnet IDs — used for the Fargate profile and as part of the cluster's subnet list."
  type        = list(string)
}

variable "public_subnet_ids" {
  description = "Public subnet IDs — included in the cluster's subnet list."
  type        = list(string)
}

variable "endpoint_public_access" {
  description = "Whether the EKS API server endpoint is reachable from the public internet."
  type        = bool
  default     = true
}

variable "endpoint_private_access" {
  description = "Whether the EKS API server endpoint is reachable from inside the VPC."
  type        = bool
  default     = true
}

variable "public_access_cidrs" {
  description = "CIDR blocks allowed to reach the public endpoint, if enabled."
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "route53_zone_arn" {
  description = "ARN of the pre-existing Route53 hosted zone ExternalDNS is allowed to write to. Not something this module manages the lifecycle of — the zone already exists outside Terraform."
  type        = string
}

variable "name_suffix" {
  description = "Optional suffix appended only to IAM role/instance-profile names, which are account-wide unique unlike the cluster name itself. Needed once this module is called more than once in the same account (e.g. a second region) with the same project_name/environment. Leave empty to preserve existing naming."
  type        = string
  default     = ""
}
