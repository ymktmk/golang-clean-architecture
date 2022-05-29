resource "aws_vpc" "vpc" {
    cidr_block           = "10.0.0.0/16"
    enable_dns_hostnames = true
    enable_dns_support   = true
    instance_tenancy     = "default"
    tags = {
        "Name" = "vpc"
    }
}

# パブリックサブネット1a
resource "aws_subnet" "public_subnet_1a" {
    availability_zone = "ap-northeast-1a"
    cidr_block        = "10.0.0.0/24"
    map_public_ip_on_launch = true
    tags = {
        "Name" = "public-subnet-1a"
    }
    vpc_id = aws_vpc.vpc.id
}

# パブリックサブネット1c
resource "aws_subnet" "public_subnet_1c" {
    availability_zone = "ap-northeast-1c"
    cidr_block        = "10.0.1.0/24"
    map_public_ip_on_launch = true
    tags = {
        "Name" = "public-subnet-1c"
    }
    vpc_id = aws_vpc.vpc.id
}

# プライベートサブネット1a
resource "aws_subnet" "private_subnet_1a" {
    availability_zone       = "ap-northeast-1a"
    cidr_block              = "10.0.2.0/24"
    tags = {
        "Name" = "private-subnet-1a"
    }
    vpc_id = aws_vpc.vpc.id
}

# プライベートサブネット1c
resource "aws_subnet" "private_subnet_1c" {
    availability_zone       = "ap-northeast-1c"
    cidr_block              = "10.0.3.0/24"
    tags = {
        "Name" = "private-subnet-1c"
    }
    vpc_id = aws_vpc.vpc.id
}

# インターネットゲートウェイ
resource "aws_internet_gateway" "igw" {
    tags = {
        "Name" = "igw"
    }
    vpc_id = aws_vpc.vpc.id
}

# ルートテーブル
resource "aws_route_table" "public_subnet_route_table" {
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.igw.id
    }
    tags = {
        "Name" = "public-subnet-route-table"
    }
    vpc_id = aws_vpc.vpc.id
}

resource "aws_route_table_association" "public_subnet_1a_association" {
    route_table_id = aws_route_table.public_subnet_route_table.id
    subnet_id      = aws_subnet.public_subnet_1a.id
}

resource "aws_route_table_association" "public_subnet_1c_association" {
    route_table_id = aws_route_table.public_subnet_route_table.id
    subnet_id      = aws_subnet.public_subnet_1c.id
}