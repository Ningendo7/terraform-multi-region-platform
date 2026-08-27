resource "aws_dynamodb_table" "this" {
  name         = "${local.name}-${var.table_name}"
  billing_mode = var.billing_mode

  hash_key  = var.hash_key
  range_key = var.range_key

  dynamic "attribute" {
    for_each = var.attributes

    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  # Global Tables (v2) require streams — only turned on when there's
  # actually a replica region to stream changes to.
  stream_enabled   = length(var.replica_regions) > 0
  stream_view_type = length(var.replica_regions) > 0 ? "NEW_AND_OLD_IMAGES" : null

  point_in_time_recovery {
    enabled = var.point_in_time_recovery_enabled
  }

  server_side_encryption {
    enabled     = true
    kms_key_arn = var.kms_key_arn
  }

  deletion_protection_enabled = var.deletion_protection_enabled

  dynamic "replica" {
    for_each = var.replica_regions

    content {
      region_name = replica.value
    }
  }

  tags = local.tags
}
