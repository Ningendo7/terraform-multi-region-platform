# EKS Module

Creates an EKS cluster with no managed node groups — capacity is Karpenter-only.

Creates:

- EKS cluster (API access-entry authentication mode; the applying principal gets a cluster-admin access entry automatically)
- Cluster IAM role
- IAM OIDC provider (IRSA — lets Kubernetes service accounts assume IAM roles directly, no static credentials)
- Node IAM role + instance profile (what Karpenter-launched EC2 instances run as)
- Karpenter controller IAM role, trust-scoped to the `system:serviceaccount:kube-system:karpenter` service account only, plus its permission policy (EC2 lifecycle actions, pricing lookup, AMI resolution, `iam:PassRole` scoped to the node role only — not `*`)
- Fargate profile for the `kube-system` namespace, so CoreDNS and the Karpenter controller itself can run before any EC2 node exists

## Inputs

| Name | Description |
|------|-------------|
| project_name | Project name |
| environment | Environment |
| kubernetes_version | EKS Kubernetes version (default `1.31`) |
| vpc_id | VPC ID the cluster deploys into |
| private_subnet_ids | Private subnet IDs (used for the Fargate profile and the cluster's subnet list) |
| public_subnet_ids | Public subnet IDs (included in the cluster's subnet list) |
| endpoint_public_access | Whether the API server is reachable from the public internet (default `true`) |
| endpoint_private_access | Whether the API server is reachable from inside the VPC (default `true`) |
| public_access_cidrs | CIDRs allowed to reach the public endpoint, if enabled (default `["0.0.0.0/0"]`) |

## Outputs

- cluster_name
- cluster_endpoint
- cluster_certificate_authority_data
- oidc_provider_arn
- karpenter_controller_role_arn
- node_iam_role_arn
- node_instance_profile_name

## Not included

Karpenter itself isn't installed by this module — only the IAM/OIDC/Fargate scaffolding it needs once it's deployed (Helm or GitOps, once ArgoCD exists). Same for any other cluster add-on.
