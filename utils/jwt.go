package utils

import (
	"gin/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJwtToken(cpf string) (string, error) {

	config.LoadEnvironmentVariables()

	var secretKey = config.GetSecretKeyApi()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cpf": cpf,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return jwt, nil

}
