provider "aws"{
    region = "ap-northeast-1"
    profile = "default"
}

terraform {
    required_version = "~> 1.1.7"
    required_providers {
        aws = {
                source  = "hashicorp/aws"
                version = "~> 4.0"
        }
    }
    # backend "s3" {
    #     bucket = "terraform-state-ymktmk"
    #     region = "ap-northeast-1"
    #     profile = "default"
    #     key = "terraform.tfstate"
    #     encrypt = true
    #     dynamodb_table = "terraform_state_lock"
    # }
}

# resource "aws_s3_bucket" "terraform_state" {
#     bucket = "terraform-state-ymktmk"
#     lifecycle {
#         prevent_destroy = true
#     }
#     versioning {
#         enabled = true
#     }
# }

# resource "aws_dynamodb_table" "terraform_state_lock" {
#     name = "terraform_state_lock"
#     read_capacity = 1
#     write_capacity = 1
#     hash_key = "LockID"

#     attribute {
#     name = "LockID"
#     type = "S"
#     }
# }