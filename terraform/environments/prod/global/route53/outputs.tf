output "hosted_zone_id" {
         description = "The ID of the Route 53 hosted zone."
         value       = data.aws_route53_zone.this.zone_id
  
}

output "domain_name" {
         description = "The domain name managed by Route 53."
         value       = data.aws_route53_zone.this.name
  
}