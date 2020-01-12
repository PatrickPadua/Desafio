provider "aws" {
  region                  = "${var.aws_region}"
}

terraform {
  backend "s3" {
    bucket = "exemplo-sqs-terraform-patrick"
    key    = "terraform.tfstate"
    region = "us-east-1"
  }
}

resource "aws_sqs_queue" "main" {
  name                        = "${var.aws_sqs_name}.fifo"
  delay_seconds             = 0
  max_message_size          = 262144
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10

  fifo_queue                  = true
  content_based_deduplication = true

  kms_master_key_id                 = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300

  tags = {
    Environment = "production"
    Provider    = "Terraform"
  }
}

resource "local_file" "sqs_url" {
    content     = "${aws_sqs_queue.main.id}"
    filename = "./output/sqs_url.hcl"
}

resource "local_file" "sqs_group" {
    content     = "${var.aws_sqs_group}"
    filename = "./output/sqs_grp.hcl"
}