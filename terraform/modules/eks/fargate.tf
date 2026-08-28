# Fargate profile: runs kube-system (CoreDNS, the Karpenter controller,
# and any other kube-system workload) without needing any EC2 node to
# exist first.

data "aws_iam_policy_document" "fargate_assume" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["eks-fargate-pods.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "fargate" {
  name               = "${local.iam_name}-eks-fargate"
  assume_role_policy = data.aws_iam_policy_document.fargate_assume.json
  tags               = local.tags
}

resource "aws_iam_role_policy_attachment" "fargate_pod_execution" {
  role       = aws_iam_role.fargate.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy"
}

resource "aws_eks_fargate_profile" "kube_system" {
  cluster_name           = aws_eks_cluster.this.name
  fargate_profile_name   = "${local.name}-kube-system"
  pod_execution_role_arn = aws_iam_role.fargate.arn
  subnet_ids             = var.private_subnet_ids

  selector {
    namespace = "kube-system"
  }
}
