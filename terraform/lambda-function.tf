resource "aws_lambda_function" "csye6225_send_verification_email" {
  filename      = "../bootstrap.zip"
  function_name = "send_verification_email"
  role          = aws_iam_role.csye6225_lambda_exec_role.arn
  handler       = "send_verification_email"
  runtime       = "provided.al2"
  timeout       = 120

  environment {
    variables = {
      SENDGRID_API_KEY  = var.sendgrid_api_key
      WEBAPP_HOSTNAME   = var.webapp_hostname
      EMAIL_SENDER_ADDR = var.email_sender_address
      EMAIL_SENDER_NAME = var.email_sender_name
    }
  }
}

resource "aws_iam_role" "csye6225_lambda_exec_role" {
  name = "lambda_exec_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_sns_rds_access" {
  role       = aws_iam_role.csye6225_lambda_exec_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
