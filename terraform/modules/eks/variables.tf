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
