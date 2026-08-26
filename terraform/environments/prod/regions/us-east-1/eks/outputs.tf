output "cluster_name" {
  description = "The name of the EKS cluster."
  value       = module.eks.cluster_name
}

output "cluster_endpoint" {
  description = "The EKS API server endpoint."
  value       = module.eks.cluster_endpoint
}

output "oidc_provider_arn" {
  description = "ARN of the cluster's IAM OIDC provider."
  value       = module.eks.oidc_provider_arn
}

output "karpenter_controller_role_arn" {
  description = "ARN of the IAM role Karpenter's service account assumes."
  value       = module.eks.karpenter_controller_role_arn
}
