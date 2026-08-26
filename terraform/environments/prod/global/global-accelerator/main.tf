resource "aws_globalaccelerator_accelerator" "this" {
  name            = var.accelerator_name
  enabled         = true
  ip_address_type = "IPV4"

  tags = {
    Name = var.accelerator_name
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