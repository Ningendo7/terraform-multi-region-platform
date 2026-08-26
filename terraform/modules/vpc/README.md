# VPC Module

Creates:

- VPC
- Internet Gateway
- Public and private subnets (one pair per AZ)
- One NAT Gateway per AZ, each with its own Elastic IP
- Public route table (shared across AZs) and one private route table per AZ
- VPC Flow Logs (CloudWatch Logs, optionally KMS-encrypted) with the IAM role/policy that delivers them

## Inputs

| Name | Description |
|------|-------------|
| project_name | Project name |
| environment | Environment |
| vpc_cidr | VPC CIDR block (default `10.0.0.0/16`) |
| az_count | Number of Availability Zones to span (default `3`) |
| enable_flow_logs | Whether to enable VPC Flow Logs (default `true`) |
| flow_log_kms_key_arn | KMS key ARN for flow log encryption (default `null`, falls back to CloudWatch's default encryption) |
| flow_log_retention_days | Flow log retention in days (default `90`) |
| public_subnet_tags | Extra tags applied only to public subnets |
| private_subnet_tags | Extra tags applied only to private subnets |

## Outputs

- vpc_id
- vpc_cidr
- public_subnet_ids
- private_subnet_ids
- availability_zones
- nat_gateway_ids

## Notes

Public and private subnets are carved from the VPC CIDR as `/20`s. Public subnets take indices 0 through `az_count - 1`; private subnets start at index 10. Indices 3-9 are deliberately left unused, reserved for a future subnet tier (e.g. database subnets) without renumbering what's already deployed.
