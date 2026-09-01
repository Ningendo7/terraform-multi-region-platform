provider "aws" {

  region = "us-east-2"

  default_tags {
    tags = {
      Project     = "terraform-multi-region-platform"
      Environment = "prod"
      Region      = "us-east-2"
      ManagedBy   = "Terraform"
    }
  }
}
