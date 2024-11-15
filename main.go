package main

import (
	"context"
	"encoding/json"
	"fmt"

	"eric-cw-hsu.io/email-sender/pkg/email_sender"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, snsEvent events.SNSEvent) error {
	for _, record := range snsEvent.Records {
		message := record.SNS.Message

		var email email_sender.VerificationEmail
		if err := json.Unmarshal([]byte(message), &email); err != nil {
			return err
		}

		if err := email_sender.SendVerificationEmail(email); err != nil {
			return err
		}

		fmt.Println("Email sent successfully")
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
