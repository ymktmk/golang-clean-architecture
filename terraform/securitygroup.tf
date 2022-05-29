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