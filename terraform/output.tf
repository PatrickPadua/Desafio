output "this_sqs_queue_url" {
  description = "The URL for the created Amazon SQS queue"
  value       = "${aws_sqs_queue.main.id}"
}

output "this_sqs_queue_group" {
  description = "The URL for the created Amazon SQS queue"
  value       = "${var.aws_sqs_group}"
}