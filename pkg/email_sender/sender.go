package email_sender

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Verification Email Template
var EMAIL_RETRY_LIMIT = 3

func SendVerificationEmail(email VerificationEmail) error {
	var err error
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	email.InitEmail()

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
