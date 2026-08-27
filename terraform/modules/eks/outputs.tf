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

output "cluster_security_group_id" {
  description = "ID of the EKS cluster's security group — the one EKS automatically manages for control-plane-to-node communication. Reference this from anything (like Aurora) that needs to allow ingress from workloads running in the cluster."
  value       = aws_eks_cluster.this.vpc_config[0].cluster_security_group_id
}

output "lb_controller_role_arn" {
  description = "ARN of the IAM role the AWS Load Balancer Controller's service account assumes."
  value       = aws_iam_role.lb_controller.arn
}

output "fargate_pod_execution_role_arn" {
  description = "ARN of the Fargate pod execution role — reusable across any additional Fargate profiles (e.g. one for argocd's namespace) so IAM doesn't need to be duplicated per profile."
  value       = aws_iam_role.fargate.arn
}