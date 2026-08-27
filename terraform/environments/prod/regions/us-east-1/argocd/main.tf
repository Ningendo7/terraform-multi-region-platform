data "terraform_remote_state" "eks" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-east-1/eks/terraform.tfstate"
    region = "us-east-1"
  }
}

data "terraform_remote_state" "vpc" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-east-1/vpc/terraform.tfstate"
    region = "us-east-1"
  }
}

# argocd needs its own Fargate profile — kube-system's profile doesn't
# cover it, and Karpenter isn't installed yet, so without this there is
# no compute path for anything in the argocd namespace at all. Reuses
# the pod execution role already built in modules/eks rather than
# duplicating that IAM setup here.
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

  # Explicit — Terraform can't infer this from the helm_release resource
  # itself, since Helm-managed pods are invisible to its dependency
  # graph. Without this, the profile and the Helm install could race,
  # reproducing the exact issue that just happened.
  depends_on = [aws_eks_fargate_profile.argocd]
}
