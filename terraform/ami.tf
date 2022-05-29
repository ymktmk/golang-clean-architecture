data "aws_ami" "ami" {
      filter {
            name   = "name"
            values = ["amzn2-ami-ecs-hvm-2.0.20201013-x86_64-ebs"]
      }
      owners = ["amazon"]
}