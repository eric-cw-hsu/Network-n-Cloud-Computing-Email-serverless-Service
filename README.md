# Serverless

An Email Sender micro-service with terraform auto deployment on AWS Lambda.

## Usage
Create `terraform.tfvars` in folder /terraform with following settings
```
sendgrid_api_key      = 
webapp_hostname       = 
email_sender_address  = 
email_sender_name     = 
```

```bash
# Deploy Service to AWS
$ make deploy-email-service

# Destroy Service
$ make destroy-email-service
```

## Folder Structure
project-root/
├── pkg/
├── main.go
├── Makefile
└── terraform/

## License
This project is licensed under a proprietary license. All rights are reserved by the owner. Unauthorized copying, distribution, or modification of this code is strictly prohibited.
