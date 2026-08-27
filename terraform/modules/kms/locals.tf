locals {

  name = var.name_suffix != "" ? "${var.project_name}-${var.environment}-${var.name_suffix}" : "${var.project_name}-${var.environment}"
}
