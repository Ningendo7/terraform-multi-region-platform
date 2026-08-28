# Karpenter controller IRSA role. Trusts only the specific Kubernetes
# service account Karpenter runs as, via the OIDC provider in main.tf —
# not "anything in the cluster."

data "aws_iam_policy_document" "karpenter_assume" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.cluster.arn]
    }

    condition {
      test     = "StringEquals"
      variable = "${local.oidc_issuer_host}:sub"
      values   = ["system:serviceaccount:kube-system:karpenter"]
    }

    condition {
      test     = "StringEquals"
      variable = "${local.oidc_issuer_host}:aud"
      values   = ["sts.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "karpenter_controller" {
  name               = "${local.iam_name}-karpenter-controller"
  assume_role_policy = data.aws_iam_policy_document.karpenter_assume.json
  tags               = local.tags
}

data "aws_iam_policy_document" "karpenter_controller" {
  statement {
    sid    = "EC2NodeManagement"
    effect = "Allow"

    actions = [
      "ec2:CreateLaunchTemplate",
      "ec2:CreateFleet",
      "ec2:CreateTags",
      "ec2:DeleteLaunchTemplate",
      "ec2:RunInstances",
      "ec2:TerminateInstances",
      "ec2:DescribeLaunchTemplates",
      "ec2:DescribeInstances",
      "ec2:DescribeSubnets",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeInstanceTypes",
      "ec2:DescribeInstanceTypeOfferings",
      "ec2:DescribeAvailabilityZones",
      "ec2:DescribeSpotPriceHistory",
      "ec2:DescribeImages",
    ]

    resources = ["*"]
  }

  statement {
    sid     = "PricingLookup"
    effect  = "Allow"
    actions = ["pricing:GetProducts"]

    resources = ["*"]
  }

  statement {
    sid     = "AMIResolution"
    effect  = "Allow"
    actions = ["ssm:GetParameter"]

    resources = ["arn:aws:ssm:*::parameter/aws/service/eks/optimized-ami/*"]
  }

  statement {
    sid     = "EKSClusterRead"
    effect  = "Allow"
    actions = ["eks:DescribeCluster"]

    resources = [aws_eks_cluster.this.arn]
  }

  statement {
    # Scoped to only the node role above — without this restriction,
    # PassRole is a privilege-escalation path (it could pass ANY role).
    sid       = "PassNodeRole"
    effect    = "Allow"
    actions   = ["iam:PassRole"]
    resources = [aws_iam_role.node.arn]
  }
}

resource "aws_iam_role_policy" "karpenter_controller" {
  name   = "${local.iam_name}-karpenter-controller"
  role   = aws_iam_role.karpenter_controller.id
  policy = data.aws_iam_policy_document.karpenter_controller.json
}
