variable "project_name" {
  description = "The name of the project."
  type        = string

}

variable "environment" {
  description = "The environment name."
  type        = string

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be 'dev', 'staging', 'prod'."
  }

}

variable "region" {
  description = "The AWS region."
  type        = string

}

variable "github_org" {
  description = "The GitHub organization or username."
  type        = string

}

variable "github_repository" {
  description = "The GitHub repository name."
  type        = string

}