resource "aws_ecr_repository" "ecr_repository" {
      name                 = "ecr_repository"
      image_tag_mutability = "MUTABLE"
      image_scanning_configuration {
            scan_on_push = true
      }
}

resource "aws_ecr_lifecycle_policy" "ecr_lifecycle_policy" {
      repository = aws_ecr_repository.ecr_repository.name

      policy = <<EOF
      {
      "rules": [
            {
                  "rulePriority": 1,
                  "description": "Delete images when count is more than 500",
                  "selection": {
                  "tagStatus": "any",
                  "countType": "imageCountMoreThan",
                  "countNumber": 500
                  },
                  "action": {
                  "type": "expire"
                  }
            }
      ]
      }
EOF
}