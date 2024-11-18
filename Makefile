build-email-service:
	@echo "Building email service..."
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap ./main.go
	zip bootstrap.zip bootstrap

.PHONY: build-email-service