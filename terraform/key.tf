# ssh-keygen -t rsa -f example -N ''
# ssh -i example ec2-user@<IP>
resource "aws_key_pair" "key_pair" {
    key_name   = "example"
    public_key = file("./example.pub")
}

resource "aws_eip" "ec2_eip" {
    vpc = true
    depends_on                = [aws_internet_gateway.igw]
}