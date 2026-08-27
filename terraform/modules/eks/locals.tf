locals {

  name = "${var.project_name}-${var.environment}"

  tags = {
    Environment = var.environment
    ManagedBy   = "Terraform"
  }

  oidc_issuer_host = replace(aws_iam_openid_connect_provider.cluster.url, "https://", "")
}
