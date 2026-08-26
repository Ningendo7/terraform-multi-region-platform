locals {

  name = "${var.project_name}-${var.environment}"

  azs = slice(data.aws_availability_zones.available.names, 0, var.az_count)

  # A /16 VPC split into /20s: public subnets take indices 0-2, private
  # start at index 10 to leave room for another subnet tier later (e.g.
  # database subnets) without renumbering what's already deployed.
  public_subnet_cidrs  = [for i in range(var.az_count) : cidrsubnet(var.vpc_cidr, 4, i)]
  private_subnet_cidrs = [for i in range(var.az_count) : cidrsubnet(var.vpc_cidr, 4, i + 10)]

  tags = {
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}
