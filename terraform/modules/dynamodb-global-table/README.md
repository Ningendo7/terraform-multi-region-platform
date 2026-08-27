# DynamoDB Global Table Module

Creates a DynamoDB table, optionally replicated across regions as a Global Table (v2 — the built-in `replica` block on `aws_dynamodb_table`, not the older standalone `aws_dynamodb_global_table` resource).

Fully generic — no application schema baked in. `hash_key`/`range_key`/`attributes` are all inputs, so this module has no opinion about what any particular table is for.

## Inputs

| Name | Description |
|------|-------------|
| project_name | Project name |
| environment | Environment |
| table_name | Logical table name (gets prefixed with project/environment) |
| billing_mode | DynamoDB billing mode (default `PAY_PER_REQUEST`) |
| hash_key | Partition key attribute name |
| range_key | Sort key attribute name (default `null`) |
| attributes | Attribute definitions for key/index attributes only — `list(object({ name = string, type = string }))` |
| replica_regions | Additional regions beyond wherever this module is applied (default `[]` — empty is a valid, non-global table) |
| kms_key_arn | KMS key for encryption at rest (default `null`, uses DynamoDB's AWS-managed key) |
| point_in_time_recovery_enabled | Default `true` |
| deletion_protection_enabled | Default `true` — a live AWS setting, not just a Terraform guard; must be explicitly disabled before the table can be destroyed |

## Outputs

- table_name
- table_arn
- table_id
- stream_arn

## Notes

`replica_regions` is the *additional* regions beyond the module's own home region (wherever its provider points) — it is not the full list of regions the table lives in. Streams (`stream_enabled`, required for Global Tables) only turn on automatically when `replica_regions` is non-empty.

No call-site wired up yet — there's no defined application driving a real table schema. This module exists so one can be added quickly once there is.
