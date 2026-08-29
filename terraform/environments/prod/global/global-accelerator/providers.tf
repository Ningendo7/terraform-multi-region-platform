# Global Accelerator's control-plane API only exists in us-west-2 —
# that's an AWS constraint, not a project choice. The accelerator and
# its listener are still global resources; this just says where to
# manage them from.
provider "aws" {

  region = "us-west-2"

  default_tags {
    tags = {
      Project     = "terraform-multi-region-platform"
      Environment = "prod"
      ManagedBy   = "Terraform"
    }
  }

}
