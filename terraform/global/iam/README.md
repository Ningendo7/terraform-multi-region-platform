# Global IAM

Terraform state responsible for global AWS identity resources.

## Purpose

This layer creates shared IAM resources used across the multi-region platform.

## Resources

- GitHub Actions OIDC identity provider
- Terraform execution IAM role
- IAM policies required for infrastructure automation

## State

This is an independent Terraform state.

State key:
global/iam/terraform.tfstate

## Dependencies

Requires:

- Terraform bootstrap completed
- S3 remote state backend
- DynamoDB state locking

## Future Consumers

Regional stacks will consume outputs from this layer:

- EKS deployments
- CI/CD workflows
- Platform automation tooling