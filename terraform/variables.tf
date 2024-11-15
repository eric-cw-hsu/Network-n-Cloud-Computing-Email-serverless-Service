variable "region" {
  description = "AWS Region"
  default     = "us-west-2"
}

variable "sendgrid_api_key" {
  description = "SendGrid API Key"
}

variable "webapp_hostname" {
  description = "Webapp Hostname"
}

variable "email_sender_address" {
  description = "Email Sender Address"
}

variable "email_sender_name" {
  description = "Email Sender Name"
}
