output "cluster_id" {
  description = "The ID of this region's Aurora cluster."
  value       = aws_rds_cluster.this.id
}

output "cluster_arn" {
  description = "The ARN of this region's Aurora cluster."
  value       = aws_rds_cluster.this.arn
}

output "cluster_endpoint" {
  description = "The writer endpoint for this region's cluster."
  value       = aws_rds_cluster.this.endpoint
}

output "cluster_reader_endpoint" {
  description = "The reader endpoint for this region's cluster."
  value       = aws_rds_cluster.this.reader_endpoint
}

output "global_cluster_id" {
  description = "The ID of the Aurora Global Database. Only set when is_primary is true — pass this into secondary-region calls as global_cluster_identifier."
  value       = var.is_primary ? aws_rds_global_cluster.this[0].id : null
}

output "security_group_id" {
  description = "ID of the security group guarding this cluster."
  value       = aws_security_group.this.id
}

output "master_user_secret_arn" {
  description = "Secrets Manager ARN holding the AWS-managed master password. Only set when is_primary is true."
  value       = var.is_primary ? aws_rds_cluster.this.master_user_secret[0].secret_arn : null
}