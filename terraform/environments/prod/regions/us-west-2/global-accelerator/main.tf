data "terraform_remote_state" "global_accelerator" {
  backend = "s3"

  config = {
    bucket = "terraform-multi-region-platform-prod-terraform-state"
    key    = "global/global-accelerator/terraform.tfstate"
    region = "us-east-1"
  }
}

# The ALB itself isn't Terraform-managed — the AWS Load Balancer
# Controller creates it from the demo-app Ingress, and its name/ARN
# aren't predictable ahead of time. Looking it up by the tags the
# controller always sets is the reliable way to find it, rather than
# guessing at a name pattern.
data "aws_resourcegroupstaggingapi_resources" "demo_app_alb" {
  resource_type_filters = ["elasticloadbalancing:loadbalancer"]

  tag_filter {
    key    = "ingress.k8s.aws/stack"
    values = ["demo/demo-app"]
  }

  tag_filter {
    key    = "elbv2.k8s.aws/cluster"
    values = ["terraform-multi-region-platform-prod"]
  }
}

resource "aws_globalaccelerator_endpoint_group" "demo_app" {
  provider = aws.global_accelerator

  listener_arn = data.terraform_remote_state.global_accelerator.outputs.listener_arn

  endpoint_group_region = "us-west-2"

  endpoint_configuration {
    endpoint_id = data.aws_resourcegroupstaggingapi_resources.demo_app_alb.resource_tag_mapping_list[0].resource_arn
    weight      = 100
  }

  health_check_protocol = "HTTP"
  health_check_path     = "/"
  health_check_port     = 80

  traffic_dial_percentage = 100
}
