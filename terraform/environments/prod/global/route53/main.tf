data "aws_route53_zone" "this" {

  zone_id = "Z02533231WT3LO3AOOYF2"

}

data "terraform_remote_state" "global_accelerator" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "global/global-accelerator/terraform.tfstate"
    region = "us-east-1"
  }
}

# The real multi-region entry point — distinct from demo.<domain> and
# demo-us-east-2.<domain>, which each region's ExternalDNS already owns
# and points straight at its own regional ALB. This one goes through
# Global Accelerator instead, so it's whichever region is actually
# healthy, not a fixed region.
resource "aws_route53_record" "app" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "app.${data.aws_route53_zone.this.name}"
  type    = "A"

  alias {
    name                   = data.terraform_remote_state.global_accelerator.outputs.dns_name
    zone_id                = data.terraform_remote_state.global_accelerator.outputs.hosted_zone_id
    evaluate_target_health = true
  }
}
