resource "aws_globalaccelerator_accelerator" "this" {
  name            = "terraform-multi-region-platform-prod"
  enabled         = true
  ip_address_type = "IPV4"

  tags = {
    Name = "terraform-multi-region-platform-prod"
  }
}

resource "aws_globalaccelerator_listener" "this" {
  accelerator_arn = aws_globalaccelerator_accelerator.this.arn
  protocol        = "TCP"
  port_range {
    from_port = 80
    to_port   = 80
  }

  port_range {
    from_port = 443
    to_port   = 443
  }
}
