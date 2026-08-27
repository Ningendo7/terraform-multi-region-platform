data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

data "aws_iam_policy_document" "key" {
  statement {
    sid    = "EnableRootAccountAccess"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }

    actions   = ["kms:*"]
    resources = ["*"]
  }

  dynamic "statement" {
    for_each = var.enable_cloudwatch_logs ? [1] : []

    content {
      sid    = "AllowCloudWatchLogs"
      effect = "Allow"

      principals {
        type        = "Service"
        identifiers = ["logs.${data.aws_region.current.region}.amazonaws.com"]
      }

      actions = [
        "kms:Encrypt*",
        "kms:Decrypt*",
        "kms:ReEncrypt*",
        "kms:GenerateDataKey*",
        "kms:Describe*",
      ]

      resources = ["*"]

      condition {
        test     = "ArnLike"
        variable = "kms:EncryptionContext:aws:logs:arn"
        values   = ["arn:aws:logs:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:log-group:*"]
      }
    }
  }
}

resource "aws_kms_key" "this" {
  description             = var.description
  enable_key_rotation     = var.enable_key_rotation
  deletion_window_in_days = var.deletion_window_in_days
  multi_region            = var.multi_region
  policy                  = data.aws_iam_policy_document.key.json

  tags = {
    Name = local.name
  }

}

resource "aws_kms_alias" "this" {
  name          = "alias/${local.name}"
  target_key_id = aws_kms_key.this.key_id
}
