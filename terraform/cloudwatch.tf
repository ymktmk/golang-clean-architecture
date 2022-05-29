# ECSのロググループ
resource "aws_cloudwatch_log_group" "log_group_for_ecs" {
    name              = "/ecs/log/go"
    retention_in_days = 7
}