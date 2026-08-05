resource "aws_kms_key" "this" {
  description         = var.description
  enable_key_rotation = var.enable_key_rotation
  deletion_window_in_days = var.deletion_window_in_days
  multi_region = var.multi_region

  tags = {
    Name = local.name
  }

}

resource "aws_kms_alias" "this" {
  name          = "alias/${local.name}"
  target_key_id = aws_kms_key.this.key_id
}