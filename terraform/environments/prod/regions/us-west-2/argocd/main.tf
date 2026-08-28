data "terraform_remote_state" "eks" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-west-2/eks/terraform.tfstate"
    region = "us-east-1"
  }
}

data "terraform_remote_state" "vpc" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-west-2/vpc/terraform.tfstate"
    region = "us-east-1"
  }
}

# argocd needs its own Fargate profile — kube-system's profile doesn't
# cover it, and Karpenter isn't installed yet, so without this there is
# no compute path for anything in the argocd namespace at all. Reuses
# the pod execution role already built in modules/eks rather than
# duplicating that IAM setup here. Fargate profile names only need to
# be unique per-cluster, so reusing the same literal as us-east-1's is
# fine — this is a completely separate cluster.
resource "aws_eks_fargate_profile" "argocd" {
  cluster_name           = data.terraform_remote_state.eks.outputs.cluster_name
  fargate_profile_name   = "terraform-multi-region-platform-prod-argocd"
  pod_execution_role_arn = data.terraform_remote_state.eks.outputs.fargate_pod_execution_role_arn
  subnet_ids             = data.terraform_remote_state.vpc.outputs.private_subnet_ids

  selector {
    namespace = "argocd"
  }
}

module "argocd" {
  source = "../../../../../modules/argocd"

  # Explicit — see us-east-1's argocd stack for why this matters.
  depends_on = [aws_eks_fargate_profile.argocd]
}
