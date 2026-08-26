output "cluster_name" {
  description = "The name of the EKS cluster."
  value       = aws_eks_cluster.this.name
}

output "cluster_endpoint" {
  description = "The EKS API server endpoint."
  value       = aws_eks_cluster.this.endpoint
}

output "cluster_certificate_authority_data" {
  description = "Base64-encoded certificate authority data, needed for kubeconfig."
  value       = aws_eks_cluster.this.certificate_authority[0].data
}

output "oidc_provider_arn" {
  description = "ARN of the cluster's IAM OIDC provider, used to build further IRSA roles."
  value       = aws_iam_openid_connect_provider.cluster.arn
}

output "karpenter_controller_role_arn" {
  description = "ARN of the IAM role Karpenter's Kubernetes service account assumes."
  value       = aws_iam_role.karpenter_controller.arn
}

output "node_iam_role_arn" {
  description = "ARN of the IAM role Karpenter-launched EC2 nodes run as."
  value       = aws_iam_role.node.arn
}

output "node_instance_profile_name" {
  description = "Name of the instance profile Karpenter attaches to nodes it launches."
  value       = aws_iam_instance_profile.node.name
}