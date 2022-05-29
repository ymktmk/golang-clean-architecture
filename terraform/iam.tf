resource "aws_iam_role" "ecs_instance_role" {
	name = "ecs_instance_role"
	assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
        "Sid": "",
        "Effect": "Allow",
        "Principal": {
            "Service": ["ec2.amazonaws.com"]
        },
        "Action": "sts:AssumeRole"
        }
    ]
}
EOF
}

# ECSタスクの実行権限
resource "aws_iam_role_policy_attachment" "ecs_task_execution_policy" {
	role = aws_iam_role.ecs_instance_role.name
	policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# ECRの読み取り権限
resource "aws_iam_role_policy_attachment" "ecr_full_access_policy" {
	role = aws_iam_role.ecs_instance_role.name
	policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}

# IAMロールをEC2インスタンスに紐つける時は、インスタンスプロファイルが必要
resource "aws_iam_role_policy_attachment" "ecs_service_policy" {
	role = aws_iam_role.ecs_instance_role.name
	policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

# ECRビルド
resource "aws_iam_role_policy_attachment" "ecs_build_policy" {
	role = aws_iam_role.ecs_instance_role.name
	policy_arn = "arn:aws:iam::aws:policy/EC2InstanceProfileForImageBuilderECRContainerBuilds"
}

# EC2インスタンスにIAMロール
resource "aws_iam_instance_profile" "ecs_instance_role" {
    name = "ecs_instance_role"
    role = aws_iam_role.ecs_instance_role.name
}