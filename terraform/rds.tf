resource "aws_db_instance" "db_instance" {
      allocated_storage      = 10
      storage_type           = "gp2"
      engine                 = "mysql"
      engine_version         = "5.7"
      instance_class         = "db.t3.micro"
      username               = "admin"
      password               = "password"
      parameter_group_name   = "default.mysql5.7"
      port                   = 3306
      # シングルAZ構成
      multi_az               = false
      db_subnet_group_name   = aws_db_subnet_group.db_subnet_group.name
      # 検証用のためfalseにする
      skip_final_snapshot    = true
      vpc_security_group_ids = [aws_security_group.rds_security_group.id]
}

# DBサブネットグループの定義
resource "aws_db_subnet_group" "db_subnet_group" {
      name       = "db_subnet_group"
      subnet_ids = [
            aws_subnet.private_subnet_1a.id,
            aws_subnet.private_subnet_1c.id
      ]
      tags = {
            Name = "db_subnet_group"
      }
}


# RDSのセキュリティーグループ
resource "aws_security_group" "rds_security_group" {
      name ="rds_security_group"
      vpc_id= aws_vpc.vpc.id
      tags = {
            Name = "rds_security_group"
      }
}

# 3306番ポート開放
resource "aws_security_group_rule" "accept3306" {
      security_group_id = aws_security_group.rds_security_group.id
      type              = "ingress"
      cidr_blocks = [aws_vpc.vpc.cidr_block]
      from_port         = 3306
      to_port           = 3306
      protocol          = "tcp"
}

resource "aws_security_group_rule" "rds_out_all" {
      security_group_id = aws_security_group.rds_security_group.id
      type              = "egress"
      cidr_blocks       = ["0.0.0.0/0"]
      from_port         = 0
      to_port           = 0
      protocol          = "-1"
}