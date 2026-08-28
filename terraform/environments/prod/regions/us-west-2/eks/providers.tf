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
