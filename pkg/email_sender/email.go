package email_sender

import (
	"fmt"
	"os"
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

func (e *VerificationEmail) InitEmail() {
	e.FromName = os.Getenv("EMAIL_SENDER_NAME")
	e.FromAddr = os.Getenv("EMAIL_SENDER_ADDR")
	e.Subject = "Email Verification"

	verifyURL := fmt.Sprintf("%s/verify?user_id=%s&token=%s", os.Getenv("WEBAPP_HOSTNAME"), e.UserId, e.Token)

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
