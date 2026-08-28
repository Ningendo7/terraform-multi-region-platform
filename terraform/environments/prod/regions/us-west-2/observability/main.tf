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

# monitoring needs its own Fargate profile — same reasoning as
# argocd's and demo-app's: its own namespace isn't covered by the
# kube-system profile.
resource "aws_eks_fargate_profile" "monitoring" {
  cluster_name           = data.terraform_remote_state.eks.outputs.cluster_name
  fargate_profile_name   = "terraform-multi-region-platform-prod-monitoring"
  pod_execution_role_arn = data.terraform_remote_state.eks.outputs.fargate_pod_execution_role_arn
  subnet_ids             = data.terraform_remote_state.vpc.outputs.private_subnet_ids

  selector {
    namespace = "monitoring"
  }
}
