provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "go_server" {
  ami           = "ami-0c55b159cbfafe1f0" # Replace with a suitable AMI ID for your region
  instance_type = "t2.micro"
  key_name       = "your-key-pair" # Replace with your EC2 key pair

  tags = {
    Name = "go-server"
  }

  user_data = <<-EOF
              #!/bin/bash
              # Update and install Docker
              apt-get update
              apt-get install -y docker.io

              # Start Docker service
              systemctl start docker
              systemctl enable docker

              # Pull and run ZincSearch Docker image
              docker run -d -p 4080:4080 --name zincsearch zincsearch/zinc:latest

              # Pull and run Go server Docker image
              docker run -d -p 8080:8080 --name go-server your-docker-repo/go-server:latest
              EOF
}

output "instance_ip" {
  value = aws_instance.go_server.public_ip
}
