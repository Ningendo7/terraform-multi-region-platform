# BOOTSTRAP

This Terraform configuration provisions the infrastructure required for Terraform itself, including:

- KMS key
- Remote state bucket
- Encryption
- State Locking
- Supporting IAM resources

This stack is deployed once before any other Terraform configuration.