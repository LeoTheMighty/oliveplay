variable "aws_region" {
  type        = string
  description = "AWS region to deploy into"
  default     = "us-east-1"
}

variable "domain_name" {
  type        = string
  description = "Root domain to reference in Route53 (e.g. catescafe.com)"
  default     = "catescafe.com"
}

variable "subdomain" {
  type        = string
  description = "Record name for the API subdomain"
  default     = "api"
}

variable "eks_cluster_name" {
  type        = string
  description = "Name of the EKS cluster"
  default     = "catescafe-eks"
}

variable "db_instance_type" {
  type        = string
  description = "RDS instance type"
  default     = "db.t3.medium"
}

variable "db_name" {
  type        = string
  description = "Database name for the app"
  default     = "catescafedb"
}

variable "db_username" {
  type        = string
  description = "Database username"
  default     = "postgres"
}

variable "db_password" {
  type        = string
  description = "Database password"
  default     = "SuperSecurePassword123!"
  sensitive   = true
}

variable "environment" {
  type        = string
  description = "Environment name (e.g. dev, prod)"
  default     = "dev"
} 