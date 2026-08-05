variable "table_name" {

         description = "DynamoDB table name."
         type = string
  
}

variable "environment" {
         description = "Environment name."
         type = string
  
}

variable "kms_key_arn" {
         description = "KMS Key ARN for server-side encryption."
         type = string
}