provider "aws" {

  region = "us-west-2"

  default_tags {
    tags = {
      Project     = "terraform-multi-region-platform"
      Environment = "prod"
      Region      = "us-west-2"
      ManagedBy   = "Terraform"
    }
  }

}

# Global Accelerator's control-plane API only exists in us-west-2 —
# this happens to be the same region as the default provider above,
# but it's kept as an explicit alias anyway so this stack matches its
# us-east-1 sibling exactly and doesn't rely on the coincidence.
provider "aws" {
  alias = "global_accelerator"

  region = "us-west-2"

  default_tags {
    tags = {
      Project     = "terraform-multi-region-platform"
      Environment = "prod"
      Region      = "us-west-2"
      ManagedBy   = "Terraform"
    }
  }

}
