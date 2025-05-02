########################################
# S3 Bucket for Worker & API
########################################

resource "aws_s3_bucket" "catescafe_storage" {
  bucket = "catescafe-${var.environment}-storage"
  acl    = "private"

  tags = {
    Environment = var.environment
    Name        = "catescafe-storage"
  }
}

########################################
# SQS Queue for Worker & App
########################################

resource "aws_sqs_queue" "catescafe_queue" {
  name = "catescafe-${var.environment}-queue"
  
  tags = {
    Environment = var.environment
  }
}

########################################
# RDS PostgreSQL
########################################

# In a real environment, you'd have a VPC and subnets. This block is only illustrative.
resource "aws_db_subnet_group" "default" {
  name       = "catescafe-${var.environment}-subnet-group"
  subnet_ids = [] # Add your real VPC subnet IDs here

  tags = {
    Name        = "catescafe-db-subnet-group"
    Environment = var.environment
  }
}

resource "aws_db_instance" "catescafe_db" {
  identifier         = "catescafe-${var.environment}-db"
  engine             = "postgres"
  engine_version     = "13"
  instance_class     = var.db_instance_type
  name               = var.db_name
  username           = var.db_username
  password           = var.db_password
  db_subnet_group_name = aws_db_subnet_group.default.name
  allocated_storage  = 20
  skip_final_snapshot = true

  # Provide your own security group(s):
  vpc_security_group_ids = [] 

  tags = {
    Environment = var.environment
  }
}

########################################
# EKS Cluster & Node Group
########################################
resource "aws_eks_cluster" "catescafe_eks" {
  name     = var.eks_cluster_name
  role_arn = ""  # Provide your EKS service role ARN

  vpc_config {
    subnet_ids = [] # Provide your subnets for EKS
  }

  tags = {
    Environment = var.environment
  }
}

resource "aws_eks_node_group" "catescafe_eks_nodes" {
  cluster_name    = aws_eks_cluster.catescafe_eks.name
  node_group_name = "catescafe-${var.environment}-nodegroup"
  node_role_arn   = "" # Provide your EC2 node IAM role ARN
  subnet_ids      = [] # Provide subnets for the node group

  scaling_config {
    desired_size = 2
    max_size     = 3
    min_size     = 1
  }

  tags = {
    Environment = var.environment
  }
}

########################################
# Kubernetes Deployments (API, Worker, Migrator)
# This is a simplified example:
########################################

resource "kubernetes_deployment" "api_deployment" {
  metadata {
    name      = "api"
    namespace = "default"
    labels = {
      app = "api"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "api"
      }
    }
    template {
      metadata {
        labels = {
          app = "api"
        }
      }
      spec {
        container {
          image = "your-registry/api-image:latest"
          name  = "api"

          env {
            name  = "SQS_QUEUE_NAME"
            value = aws_sqs_queue.catescafe_queue.name
          }
          env {
            name  = "S3_BUCKET_NAME"
            value = aws_s3_bucket.catescafe_storage.bucket
          }
          env {
            name  = "DB_HOST"
            value = aws_db_instance.catescafe_db.address
          }
          env {
            name  = "DB_NAME"
            value = var.db_name
          }
          env {
            name  = "DB_USER"
            value = var.db_username
          }
          env {
            name  = "DB_PASS"
            value = var.db_password
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment" "worker_deployment" {
  metadata {
    name      = "worker"
    namespace = "default"
    labels = {
      app = "worker"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "worker"
      }
    }
    template {
      metadata {
        labels = {
          app = "worker"
        }
      }
      spec {
        container {
          image = "your-registry/worker-image:latest"
          name  = "worker"

          env {
            name  = "SQS_QUEUE_NAME"
            value = aws_sqs_queue.catescafe_queue.name
          }
          env {
            name  = "S3_BUCKET_NAME"
            value = aws_s3_bucket.catescafe_storage.bucket
          }
          env {
            name  = "DB_HOST"
            value = aws_db_instance.catescafe_db.address
          }
          env {
            name  = "DB_NAME"
            value = var.db_name
          }
          env {
            name  = "DB_USER"
            value = var.db_username
          }
          env {
            name  = "DB_PASS"
            value = var.db_password
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment" "migrator_deployment" {
  metadata {
    name      = "migrator"
    namespace = "default"
    labels = {
      app = "migrator"
    }
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "migrator"
      }
    }
    template {
      metadata {
        labels = {
          app = "migrator"
        }
      }
      spec {
        container {
          image = "your-registry/migrator-image:latest"
          name  = "migrator"

          env {
            name  = "DB_HOST"
            value = aws_db_instance.catescafe_db.address
          }
          env {
            name  = "DB_NAME"
            value = var.db_name
          }
          env {
            name  = "DB_USER"
            value = var.db_username
          }
          env {
            name  = "DB_PASS"
            value = var.db_password
          }
        }
      }
    }
  }
} 