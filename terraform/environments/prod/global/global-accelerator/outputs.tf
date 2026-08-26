output "accelerator_arn" {
  description = "The ARN of the Global Accelerator."
  value       = aws_globalaccelerator_accelerator.this.arn

}

output "listener_arn" {
  description = "The ARN of the Global Accelerator listener."
  value       = aws_globalaccelerator_listener.this.arn
}