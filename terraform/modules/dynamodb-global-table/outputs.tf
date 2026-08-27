output "table_name" {
  description = "The name of the DynamoDB table."
  value       = aws_dynamodb_table.this.name
}

output "table_arn" {
  description = "The ARN of the DynamoDB table."
  value       = aws_dynamodb_table.this.arn
}

output "table_id" {
  description = "The ID of the DynamoDB table."
  value       = aws_dynamodb_table.this.id
}

output "stream_arn" {
  description = "The ARN of the table's DynamoDB Stream, if enabled."
  value       = aws_dynamodb_table.this.stream_arn
}