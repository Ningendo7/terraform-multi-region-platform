resource "aws_rds_global_cluster" "this" {
  count = var.is_primary ? 1 : 0

  global_cluster_identifier = "${local.name}-${var.identifier}"
  engine                    = "aurora-postgresql"
  engine_version            = var.engine_version
  database_name             = var.database_name
  storage_encrypted         = true
  deletion_protection       = var.deletion_protection
}

resource "aws_db_subnet_group" "this" {
  name       = "${local.name}-${var.identifier}"
  subnet_ids = var.subnet_ids

  tags = local.tags
}

resource "aws_security_group" "this" {
  name        = "${local.name}-${var.identifier}"
  description = "Aurora ${var.identifier} — allows Postgres only from allow-listed security groups"
  vpc_id      = var.vpc_id

  tags = local.tags
}

resource "aws_vpc_security_group_ingress_rule" "this" {
  for_each = toset(var.allowed_security_group_ids)

  security_group_id            = aws_security_group.this.id
  referenced_security_group_id = each.value
  from_port                    = 5432
  to_port                      = 5432
  ip_protocol                  = "tcp"
}

resource "aws_vpc_security_group_egress_rule" "all" {
  security_group_id = aws_security_group.this.id
  ip_protocol       = "-1"
  cidr_ipv4         = "0.0.0.0/0"
}

resource "aws_rds_cluster" "this" {
  cluster_identifier = "${local.name}-${var.identifier}"

  engine         = "aurora-postgresql"
  engine_version = var.engine_version
  engine_mode    = "provisioned"

  global_cluster_identifier = var.is_primary ? aws_rds_global_cluster.this[0].id : var.global_cluster_identifier

  # Only the cluster that originates the global database sets these —
  # secondary regions replicate from it and must leave them unset.
  database_name               = var.is_primary ? var.database_name : null
  master_username             = var.is_primary ? var.master_username : null
  manage_master_user_password = var.is_primary ? true : null

  db_subnet_group_name   = aws_db_subnet_group.this.name
  vpc_security_group_ids = [aws_security_group.this.id]

  storage_encrypted = true
  kms_key_id        = var.kms_key_arn

  iam_database_authentication_enabled = var.iam_database_authentication_enabled

  backup_retention_period   = var.backup_retention_period
  deletion_protection       = var.deletion_protection
  skip_final_snapshot       = var.skip_final_snapshot
  final_snapshot_identifier = var.skip_final_snapshot ? null : "${local.name}-${var.identifier}-final"

  tags = local.tags
}

resource "aws_rds_cluster_instance" "this" {
  count = var.instance_count

  identifier         = "${local.name}-${var.identifier}-${count.index}"
  cluster_identifier = aws_rds_cluster.this.id

  engine         = aws_rds_cluster.this.engine
  engine_version = aws_rds_cluster.this.engine_version

  instance_class = var.instance_class

  db_subnet_group_name = aws_db_subnet_group.this.name

  tags = local.tags
}