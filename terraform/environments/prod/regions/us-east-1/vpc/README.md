# Region: us-east-1 — VPC

Terraform state responsible for the VPC in us-east-1.

## Purpose

Calls `modules/vpc` with region-specific inputs. Region and environment are literals here, not variables — this directory *is* `us-east-1`/`prod`, so nothing can point it at the wrong place by mistake.

## Resources

See `modules/vpc` for what's actually created — this stack is a thin call-site, not where the logic lives.

## State

This is an independent Terraform state.

State key:
regions/us-east-1/vpc/terraform.tfstate

## Dependencies

Requires:

- Terraform bootstrap completed. Reads its KMS key ARN via `terraform_remote_state` (bootstrap's local state) to encrypt VPC Flow Logs — this is optional and wrapped in `try(...)`, so it degrades gracefully to CloudWatch's default encryption if bootstrap hasn't been applied yet, and self-upgrades to the real key with no code change once it has.

## Future Consumers

- `regions/us-east-1/eks` reads this stack's `vpc_id` and subnet ID outputs via `terraform_remote_state`, against the real S3 backend.
