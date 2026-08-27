# Aurora Global Database Module

Creates an Aurora PostgreSQL cluster in one region, optionally as part of a multi-region Aurora Global Database.

Unlike `dynamodb-global-table` (where every replica is a sub-block of one resource), Aurora Global Database requires a **separate `aws_rds_cluster` per region**, each created from that region's own provider. This module handles one region per call, in one of two modes:

- `is_primary = true` (default) — creates the `aws_rds_global_cluster` plus the first regional cluster. Sets the master username, initial database name, and AWS-managed master password.
- `is_primary = false` — creates a regional cluster that attaches to an **existing** global cluster (`global_cluster_identifier`, taken from the primary call's `global_cluster_id` output). Does not set master credentials or database name — those only exist on the originating cluster; secondary regions replicate them.

The module doesn't handle cross-region provider aliasing itself — same as every other stack in this repo, that's the call-site's job: a `us-west-2` call-site (once that region exists) would just use its own regional provider and call this module with `is_primary = false`.

## Inputs

| Name | Description |
|------|-------------|
| project_name | Project name |
| environment | Environment |
| identifier | Logical database name (gets prefixed with project/environment) |
| is_primary | `true` creates the global cluster + primary regional cluster; `false` attaches a secondary (default `true`) |
| global_cluster_identifier | Existing global cluster ID to attach to — required when `is_primary` is `false` |
| engine_version | Aurora PostgreSQL engine version (default `16.4`) |
| database_name | Initial database name — primary only |
| master_username | Master username — primary only (default `app_admin`) |
| instance_class | Instance class for every instance in this region (default `db.r6g.large`) |
| instance_count | Number of instances (writer + readers) in this region (default `1`) |
| vpc_id | VPC ID this region's cluster deploys into |
| subnet_ids | Private subnet IDs for the DB subnet group |
| allowed_security_group_ids | Security groups allowed to reach Postgres (e.g. the EKS cluster's security group) |
| kms_key_arn | KMS key for storage encryption (default `null`, uses RDS's AWS-managed key) |
| backup_retention_period | Backup retention in days (default `7`) |
| deletion_protection | Live AWS setting; must be explicitly disabled before destroy (default `true`) |
| skip_final_snapshot | Whether to skip a final snapshot on destroy (default `false`) |
| iam_database_authentication_enabled | Allows IRSA-derived IAM credentials to authenticate alongside the master password (default `true`) |

## Outputs

- cluster_id
- cluster_arn
- cluster_endpoint
- cluster_reader_endpoint
- global_cluster_id (primary only)
- security_group_id
- master_user_secret_arn (primary only — the Secrets Manager ARN holding the AWS-managed master password; the actual password is never in Terraform state)

## Notes

Master credentials use RDS-managed master passwords (`manage_master_user_password = true`), not a plaintext variable — the real secret lives in Secrets Manager, never in Terraform state.

Security group ingress is built from the modern `aws_vpc_security_group_ingress_rule`/`egress_rule` resources rather than inline `ingress {}`/`egress {}` blocks on `aws_security_group`, so adding or removing an allow-listed security group doesn't force replacement of the whole security group.

No call-site wired up yet — same reasoning as `dynamodb-global-table`: there's no defined application/schema to justify one. Add one once there's a real workload.
