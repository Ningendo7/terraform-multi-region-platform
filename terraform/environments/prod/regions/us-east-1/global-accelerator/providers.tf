provider "aws" {

  region = "us-east-1"

  default_tags {
    tags = {
      Project     = "terraform-multi-region-platform"
      Environment = "prod"
      Region      = "us-east-1"
      ManagedBy   = "Terraform"
    }
  }

}

# Global Accelerator's control-plane API only exists in us-west-2 —
# that applies to this endpoint group too, even though it points at a
# us-east-1 ALB. This alias is used only for the endpoint group itself;
# the ALB lookup below still runs against the default (us-east-1)
# provider, since that's actually where the ALB lives.
provider "aws" {
  alias = "global_accelerator"

  region = "us-west-2"

  default_tags {
    tags = {
      Project     = "terraform-multi-region-platform"
      Environment = "prod"
      Region      = "us-east-1"
      ManagedBy   = "Terraform"
    }
  }

}
