variable "project_name" {
  description = "The name of the project"
  type        = string

}

variable "environment" {
  description = "The environment name"
  type        = string

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be 'dev', 'staging', 'prod'."
  }

}

variable "region" {
  description = "The AWS region"
  type        = string

}

variable "hosted_zone_id" {
  description = "The ID of the Existing Route 53 hosted zone."
  type        = string

}

variable "domain_name" {
  description = "Domain managed by this platform."
  type        = string

}