build-email-service:
	@echo "Building email service..."
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap ./main.go
	zip bootstrap.zip bootstrap

apply-email-service:
	@echo "Deploying email service..."
	cd terraform && terraform init && terraform apply -auto-approve

deploy-email-service: build-email-service destroy-email-service apply-email-service	

destroy-email-service:
	@echo "Destroying email service..."
	cd terraform && terraform destroy -auto-approve

.PHONY: build-email-service deploy-email-service destroy-email-service