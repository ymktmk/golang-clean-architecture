# AutoScaling EC2起動設定
resource "aws_launch_configuration" "ecs_launch_config" {
      name                 = "cluster"
      image_id             = data.aws_ami.ami.id
      iam_instance_profile = aws_iam_instance_profile.ecs_instance_role.name
      security_groups      = [aws_security_group.security_group.id]
      user_data            = <<EOF
#!/bin/bash
echo ECS_CLUSTER=online-code-cluster >> /etc/ecs/ecs.config;
INSTANCE_ID=`curl http://169.254.169.254/latest/meta-data/instance-id`
aws ec2 associate-address --allocation-id ${aws_eip.ec2_eip.allocation_id} --instance $INSTANCE_ID
EOF
      instance_type        = "t2.micro"
      key_name               = aws_key_pair.key_pair.id
      lifecycle {
            create_before_destroy = true
      }
}

# AutoScaling Group
resource "aws_autoscaling_group" "ecs_autoscaling_group" {
      name                      = "ecs_autoscaling_group"
      # 1aに配置する
      vpc_zone_identifier       = [aws_subnet.public_subnet_1a.id]
      launch_configuration      = aws_launch_configuration.ecs_launch_config.name
      min_size                  = 0
      max_size                  = 1
      desired_capacity          = 1
      health_check_grace_period = 0
      protect_from_scale_in = true
      lifecycle {
            create_before_destroy = true
      }
      health_check_type         = "EC2"
}

# Security Group
resource "aws_security_group" "security_group" {
      name = "security_group"
      vpc_id = aws_vpc.vpc.id
      tags = {
            Name = "security_group"
      }
}

# 9000番ポート開放
resource "aws_security_group_rule" "accept9000" {
      security_group_id = aws_security_group.security_group.id
      type = "ingress"
      cidr_blocks = ["0.0.0.0/0"]
      from_port = 9000
      to_port = 9000
      protocol = "tcp"
}

# 22番ポート開放
resource "aws_security_group_rule" "accept22" {
      security_group_id = aws_security_group.security_group.id
      type              = "ingress"
      cidr_blocks       = ["0.0.0.0/0"]
      from_port         = 22
      to_port           = 22
      protocol          = "tcp"
}

# アウトバウンドルール
resource "aws_security_group_rule" "out_all" {
      security_group_id = aws_security_group.security_group.id
      type              = "egress"
      cidr_blocks       = ["0.0.0.0/0"]
      from_port         = 0
      to_port           = 0
      protocol          = "-1"
}