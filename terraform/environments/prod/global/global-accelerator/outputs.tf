output "accelerator_arn" {
  description = "The ARN of the Global Accelerator."
  value       = aws_globalaccelerator_accelerator.this.arn

}

output "listener_arn" {
  description = "The ARN of the Global Accelerator listener."
  value       = aws_globalaccelerator_listener.this.arn
}

output "dns_name" {
  description = "The accelerator's static DNS name — what Route53 aliases to."
  value       = aws_globalaccelerator_accelerator.this.dns_name
}

output "hosted_zone_id" {
  description = "The accelerator's hosted zone ID, needed for the Route53 alias record. Fixed per-service by AWS, not account-specific — reading it off the resource avoids hardcoding that constant."
  value       = aws_globalaccelerator_accelerator.this.hosted_zone_id
}