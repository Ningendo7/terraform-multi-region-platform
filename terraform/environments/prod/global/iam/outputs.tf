output "terraform_role_arn" {
  description = "Terraform execution role ARN."
  value       = aws_iam_role.terraform.arn
}

output "github_oidc_provider_arn" {
  description = "GitHub OIDC provider ARN."
  value       = aws_iam_openid_connect_provider.github.arn
}