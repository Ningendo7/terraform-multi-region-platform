# Region: us-east-1 — EKS

Terraform state responsible for the EKS cluster in us-east-1.

## Purpose

Calls `modules/eks` with region-specific inputs, sourced from the `vpc` stack's real state.

## Resources

See `modules/eks` for what's actually created — this stack is a thin call-site, not where the logic lives.

## State

This is an independent Terraform state.

State key:
regions/us-east-1/eks/terraform.tfstate

## Dependencies

Requires:

- `regions/us-east-1/vpc` applied. This stack reads its VPC ID and subnet IDs via `terraform_remote_state` against the real S3 backend — and unlike bootstrap's KMS key, these are **not** optional: an EKS cluster genuinely cannot exist without a VPC, so this stack's `plan` will fail until `vpc` is actually applied. That's expected behavior, not a bug — there's no meaningful default to fall back to.

## Future Consumers

- Karpenter, ArgoCD/GitOps, and cluster add-ons get installed against this cluster once it exists — via Helm/GitOps, not Terraform.
