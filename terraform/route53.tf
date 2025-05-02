########################################
# Route53 Zone & Record for api.catescafe.com
########################################

# If the domain is already hosted in Route53, use a data lookup. 
# Otherwise, uncomment the resource block and manage it in your account.

# data "aws_route53_zone" "root_zone" {
#   name         = var.domain_name
#   private_zone = false
# }

# Alternatively, if you own the domain in R53 already, just reference it:
resource "aws_route53_zone" "root_zone" {
  name = var.domain_name
}

########################################
# Create an ACM certificate for the subdomain
########################################
resource "aws_acm_certificate" "api_cert" {
  domain_name       = "${var.subdomain}.${var.domain_name}"
  validation_method = "DNS"

  tags = {
    Environment = var.environment
    Name        = "api-certificate"
  }
}

resource "aws_route53_record" "api_cert_validation" {
  zone_id = aws_route53_zone.root_zone.zone_id
  name    = aws_acm_certificate.api_cert.domain_validation_options[0].resource_record_name
  type    = aws_acm_certificate.api_cert.domain_validation_options[0].resource_record_type
  records = [aws_acm_certificate.api_cert.domain_validation_options[0].resource_record_value]
  ttl     = 300
}

resource "aws_acm_certificate_validation" "api_cert_validation" {
  certificate_arn         = aws_acm_certificate.api_cert.arn
  validation_record_fqdns = [aws_route53_record.api_cert_validation.fqdn]
}

########################################
# API Gateway with custom domain
########################################
resource "aws_api_gateway_rest_api" "catescafe_api" {
  name        = "catescafe-api"
  description = "API Gateway for catescafe EKS service"
}

# Example root resource
resource "aws_api_gateway_resource" "root" {
  rest_api_id = aws_api_gateway_rest_api.catescafe_api.id
  parent_id   = aws_api_gateway_rest_api.catescafe_api.root_resource_id
  path_part   = "v1"
}

resource "aws_api_gateway_method" "v1_get" {
  rest_api_id   = aws_api_gateway_rest_api.catescafe_api.id
  resource_id   = aws_api_gateway_resource.root.id
  http_method   = "GET"
  authorization = "NONE"
}

# Placeholder integration responding to /v1 GET (In practice, 
# you might set up a VPC Link or an HTTP integration to your EKS service/ALB)
resource "aws_api_gateway_integration" "v1_get_integration" {
  rest_api_id             = aws_api_gateway_rest_api.catescafe_api.id
  resource_id             = aws_api_gateway_resource.root.id
  http_method             = aws_api_gateway_method.v1_get.http_method
  integration_http_method = "GET"
  type                    = "MOCK"
}

resource "aws_api_gateway_deployment" "catescafe_api_deployment" {
  depends_on = [aws_api_gateway_integration.v1_get_integration]
  rest_api_id = aws_api_gateway_rest_api.catescafe_api.id
  stage_name  = var.environment
}

# Create custom domain pointing to API Gateway
resource "aws_api_gateway_domain_name" "custom_domain" {
  domain_name = "${var.subdomain}.${var.domain_name}"
  certificate_arn = aws_acm_certificate_validation.api_cert_validation.certificate_arn
}

resource "aws_api_gateway_base_path_mapping" "api_base_path" {
  domain_name = aws_api_gateway_domain_name.custom_domain.domain_name
  rest_api_id = aws_api_gateway_rest_api.catescafe_api.id
  stage_name  = var.environment
}

# Route53 record for the custom domain to point to the API Gateway
resource "aws_route53_record" "api_gateway_alias" {
  zone_id = aws_route53_zone.root_zone.zone_id
  name    = "${var.subdomain}.${var.domain_name}"
  type    = "A"
  alias {
    name                   = aws_api_gateway_domain_name.custom_domain.cloudfront_domain_name
    zone_id                = aws_api_gateway_domain_name.custom_domain.cloudfront_zone_id
    evaluate_target_health = true
  }
} 