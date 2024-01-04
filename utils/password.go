package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"github.com/aws/aws-sdk-go/aws"
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

func enviarEmail(destinatario, mensagem string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}

	sesClient := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(destinatario)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(mensagem),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Recuperação de Senha"),
			},
		},
		Source: aws.String("seu-email@dominio.com"),
	}

	_, err = sesClient.SendEmail(input)
	return err
}
