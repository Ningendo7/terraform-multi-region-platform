locals {

  name = "${var.project_name}-${var.environment}"

  # Only for IAM role/instance-profile names — those are account-wide
  # unique, unlike the cluster name itself (EKS cluster names only need
  # to be unique per region+account, so two regions can both be named
  # "terraform-multi-region-platform-prod" without conflict). Needed
  # once this module is called more than once in the same account
  # (e.g. a second region) with the same project_name/environment.
  # Leave name_suffix empty to preserve existing naming.
  iam_name = var.name_suffix != "" ? "${local.name}-${var.name_suffix}" : local.name

  tags = {
    Environment = var.environment
    ManagedBy   = "Terraform"
  }

  oidc_issuer_host = replace(aws_iam_openid_connect_provider.cluster.url, "https://", "")
}
