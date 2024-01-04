package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"gin/config"
	"math/big"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkPasswordHash(password, hash string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil)) == hash
}

func GenerateRandomToken() (string, error) {
	tokenLength := 6

	allowedChars := "0123456789"

	tokenBytes := make([]byte, tokenLength)
	for i := 0; i < tokenLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedChars))))
		if err != nil {
			return "", err
		}
		tokenBytes[i] = allowedChars[randomIndex.Int64()]
	}

	token := string(tokenBytes)

	return token, nil
}

func SendEmailAwsSes(subject, body, recipient string) error {

	config.LoadEnvironmentVariables()

	var (
		region       = config.GetRegionAws()
		awsAccessKey = config.GetAwsAccessKey()
		awsSecretKey = config.GetAwsSecretKey()
		awsTokenKey  = config.GetAwsTokenKey()
		emailSource  = config.GetAwsEmailSource()
	)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			awsAccessKey,
			awsSecretKey,
			awsTokenKey),
	})

	if err != nil {
		return err
	}

	svc := ses.New(sess)

	emailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(emailSource),
	}

	_, err = svc.SendEmail(emailInput)

	return err
}
