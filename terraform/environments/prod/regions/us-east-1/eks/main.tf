# vpc uses the real S3 backend (unlike bootstrap's local state), so this
# reads the same bucket/key platformctl computed for it.
data "terraform_remote_state" "vpc" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "regions/us-east-1/vpc/terraform.tfstate"
    region = "us-east-1"
  }
}

module "eks" {
  source = "../../../../../modules/eks"

  project_name = "terraform-multi-region-platform"
  environment  = "prod"

  # No try()/fallback here — confirmed this actually matters, not just
  # in theory: unlike bootstrap's local-state read (empty-but-valid
  # when unapplied), the S3-backed remote_state above hard-errors with
  # "Unable to find remote state" when vpc hasn't been applied yet.
  # try() wraps the *attribute access*, but this failure happens at the
  # data source read itself, before any attribute is touched — so it
  # couldn't have been rescued that way even if we wanted to. The real
  # fix is just: apply vpc first. That's the correct dependency order,
  # not a workaround to design around.
  vpc_id             = data.terraform_remote_state.vpc.outputs.vpc_id
  private_subnet_ids = data.terraform_remote_state.vpc.outputs.private_subnet_ids
  public_subnet_ids  = data.terraform_remote_state.vpc.outputs.public_subnet_ids

  # Pre-existing zone, not managed by this Terraform project — stable
  # across destroy/recreate cycles here, unlike vpc_id above.
  route53_zone_arn = "arn:aws:route53:::hostedzone/Z02533231WT3LO3AOOYF2"
}
