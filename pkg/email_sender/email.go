package email_sender

import (
	"fmt"
)

type VerificationEmail struct {
	FromName string
	FromAddr string
	ToName   string `json:"to_name"`
	ToAddr   string `json:"to_addr"`

	UserId string `json:"user_id"`
	Token  string `json:"token"`

	Subject          string
	plainTextContent string
	htmlContent      string
}

func (e *VerificationEmail) InitEmail(secretValues *Secrets) {
	e.FromName = secretValues.EmailSenderName
	e.FromAddr = secretValues.EmailSenderAddr
	e.Subject = "Email Verification"

	verifyURL := fmt.Sprintf("%s/verify?user_id=%s&token=%s", secretValues.WebAppHostname, e.UserId, e.Token)

	e.plainTextContent = fmt.Sprintf(`
		Hi %s,

		Please verify your email address by clicking the link below.

		%s
	`, e.ToName, verifyURL)

	e.htmlContent = fmt.Sprintf(`
		<html>
			<body>
				<p>Hi %s,</p>

				<p>Please verify your email address by clicking the link below.</p>

				<a href="%s">%s</a>
			</body>
		</html>
	`, e.ToName, verifyURL, verifyURL)
}
