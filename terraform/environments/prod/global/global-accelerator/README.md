# Global Accelerator

Terraform state responsible for global traffic distribution for the multi-region platform.

## Purpose

This layer creates the global entry point for an active-active multi-region architecture.
AWS Global Accelerator provides a static anycast IP entry point and routes traffic to healthy regional endpoints.

## Architecture

Traffic flow:Users
|
v
Route53
|
v
AWS Global Accelerator
|
+----------------------+
| |
v v
us-east-1 eu-west-1EKS workload EKS workload

Each region runs identical workloads.


Global Accelerator provides:
- Global entry point- Health-based routing- Regional endpoint distribution- Improved network performance

## Resources

This state manages:
- Global Accelerator- Accelerator listener- Global traffic configuration
Regional endpoint attachments are managed by regional infrastructure states.

## Ownership Model

This state owns global routing resources.
It does not create:
- VPCs- Load balancers- EKS clusters- Regional networking
Regional stacks register their endpoints with the accelerator.

## State

This is an independent Terraform state.
Example:

global/global-accelerator/terraform.tfstate

## Dependencies
Requires:
- Terraform bootstrap backend- Regional workload endpoints
## Outputs
Provides:
- Accelerator ARN- Listener ARN- Global entry point information
## Deployment
Backend configuration is supplied during Terraform initialization.
Example:
```bashterraform init \  -backend-config=<backend-config-file>