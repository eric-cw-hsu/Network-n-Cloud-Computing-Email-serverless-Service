package email_sender

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Verification Email Template
var EMAIL_RETRY_LIMIT = 3

func SendVerificationEmail(email VerificationEmail, secretValues *Secrets) error {
	var err error
	client := sendgrid.NewSendClient(secretValues.SendGridAPIKey)

	email.InitEmail(secretValues)

	for i := 0; i < EMAIL_RETRY_LIMIT; i++ {
		res, err := client.Send(mail.NewSingleEmail(
			mail.NewEmail(email.FromName, email.FromAddr),
			email.Subject,
			mail.NewEmail(email.ToName, email.ToAddr),
			email.plainTextContent,
			email.htmlContent,
		))

		log.Printf("Email sent to %s with status code %d, body: %s", email.ToAddr, res.StatusCode, res.Body)

		if err == nil {
			break
		}
	}

	return err
}
