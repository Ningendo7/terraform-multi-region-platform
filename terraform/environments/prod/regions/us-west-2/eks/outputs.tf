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

output "cluster_security_group_id" {
  description = "ID of the EKS cluster's security group, for anything (like Aurora) that needs to allow ingress from workloads running in the cluster."
  value       = module.eks.cluster_security_group_id
}

output "lb_controller_role_arn" {
  description = "ARN of the IAM role the AWS Load Balancer Controller's service account assumes."
  value       = module.eks.lb_controller_role_arn
}

output "cluster_certificate_authority_data" {
  description = "Base64-encoded cluster CA certificate, needed to authenticate to the cluster (e.g. from the argocd stack)."
  value       = module.eks.cluster_certificate_authority_data
}

output "fargate_pod_execution_role_arn" {
  description = "ARN of the Fargate pod execution role, reusable for additional Fargate profiles."
  value       = module.eks.fargate_pod_execution_role_arn
}

output "external_dns_role_arn" {
  description = "ARN of the IAM role ExternalDNS's service account assumes."
  value       = module.eks.external_dns_role_arn
}
