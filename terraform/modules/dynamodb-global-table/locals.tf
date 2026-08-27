locals {

  name = "${var.project_name}-${var.environment}"

  tags = {
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}