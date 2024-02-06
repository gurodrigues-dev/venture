package utils

import (
	"fmt"
	"gin/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func checkTypeDocument(document *string) string {

	switch len(*document) {
	case 11:
		return "CPF"
	case 14:
		return "CNPJ"
	default:
		return "invalid"
	}

}

func CreateJwtToken(document string) (string, error) {

	typeOfDocument := checkTypeDocument(&document)

	config.LoadEnvironmentVariables()

	var secretKey = config.GetSecretKeyApi()

	if typeOfDocument == "CPF" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"cpf": document,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		jwt, err := token.SignedString([]byte(secretKey))

		if err != nil {
			return "", err
		}

		return jwt, nil
	}

	if typeOfDocument == "CNPJ" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"cnpj": document,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

		jwt, err := token.SignedString([]byte(secretKey))

		if err != nil {
			return "", err
		}

		return jwt, nil
	}

	return "", fmt.Errorf("The type of document is invalid")
}
