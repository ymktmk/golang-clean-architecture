output "elastic_ip" {
    description = "The EC2 Instance Elastic IP..."
    value = aws_eip.ec2_eip.public_ip
}

output "rds_endpoint" {
    description = "The RDS Endpoint..."
    value       = aws_db_instance.db_instance.endpoint
}