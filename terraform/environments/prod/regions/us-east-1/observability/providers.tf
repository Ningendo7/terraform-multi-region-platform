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
