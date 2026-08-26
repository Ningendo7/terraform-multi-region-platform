output "vpc_id" {
  description = "The ID of the VPC."
  value       = aws_vpc.this.id
}

output "vpc_cidr" {
  description = "The CIDR block of the VPC."
  value       = aws_vpc.this.cidr_block
}

output "public_subnet_ids" {
  description = "IDs of the public subnets, one per AZ."
  value       = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  description = "IDs of the private subnets, one per AZ."
  value       = aws_subnet.private[*].id
}

output "availability_zones" {
  description = "The Availability Zones the VPC's subnets are spread across."
  value       = local.azs
}

output "nat_gateway_ids" {
  description = "IDs of the NAT gateways, one per AZ."
  value       = aws_nat_gateway.this[*].id
}
