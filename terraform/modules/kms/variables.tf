variable "project_name" {
  description = "The name of the project."
  type        = string

}

variable "environment" {
  description = "The environment name."
  type        = string

}

variable "description" {
  description = "The description of the KMS key."
  type        = string

}

variable "enable_key_rotation" {
  description = "Enable automatic key rotation."
  type        = bool
  default     = true

}

variable "deletion_window_in_days" {
  description = "The waiting period before the KMS key is deleted."
  type        = number
  default     = 30

}

variable "multi_region" {
  description = "Whether the KMS key is multi-region."
  type        = bool
  default     = false

}