package email_sender

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Secrets struct {
	SendGridAPIKey  string `json:"SENDGRID_API_KEY"`
	WebAppHostname  string `json:"WEBAPP_HOSTNAME"`
	EmailSenderAddr string `json:"EMAIL_SENDER_ADDR"`
	EmailSenderName string `json:"EMAIL_SENDER_NAME"`
}

func GetSecrets(ctx context.Context) (*Secrets, error) {
	secretArn := os.Getenv("SECRET_ARN")
	if secretArn == "" {
		return nil, fmt.Errorf("SECRET_ARN environment variable not set")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	secretOutput, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretArn),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve secret: %w", err)
	}

	var secrets Secrets
	err = json.Unmarshal([]byte(*secretOutput.SecretString), &secrets)
	if err != nil {
		return nil, fmt.Errorf("unable to parse secret JSON: %w", err)
	}

	return &secrets, nil
}
