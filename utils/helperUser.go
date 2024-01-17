package utils

import (
	"fmt"
	"gin/models"
	"gin/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserAndAdressFromRequest(c *gin.Context) (*models.CreateUser, *models.Endereco) {

	hashedPassword := HashPassword(c.PostForm("password"))

	user := &models.CreateUser{
		Name:     c.PostForm("name"),
		Password: hashedPassword,
		CPF:      c.PostForm("cpf"),
		RG:       c.PostForm("rg"),
		ID:       uuid.New(),
		Email:    c.PostForm("email"),
	}

	endereco := &models.Endereco{
		Rua:         c.PostForm("rua"),
		Numero:      c.PostForm("numero"),
		Complemento: c.PostForm("complemento"),
		Cidade:      c.PostForm("cidade"),
		Estado:      c.PostForm("estado"),
		CEP:         c.PostForm("cep"),
	}

	return user, endereco

}

func VerifyUserAndPassword(c *gin.Context) (bool, error) {

	cpf := c.PostForm("cpf")

	hashedPassword := HashPassword(c.PostForm("password"))

	match, err := repository.VerifyPasswordByCpf(cpf, hashedPassword)

	return match, err

}

func VerifyCpf(c *gin.Context) (bool, error) {

	cpfJwtToken, found := c.Get("cpf")

	fmt.Println(cpfJwtToken)

	if !found {
		return false, fmt.Errorf("Erro ao encontrar Token")
	}

	cpfRequest := c.Param("cpf")

	cpfMatch := cpfJwtToken == cpfRequest

	return cpfMatch, nil
}

func ValidateDocsUser(user *models.CreateUser, endereco *models.Endereco) (bool, string) {

	validateCPF := IsCPF(user.CPF)

	if !validateCPF {

		return false, "cpf invalid, type and try again."

	}

	validateCEP := IsCEP(endereco.CEP)

	if !validateCEP {

		return false, "cep invalid, type and try again."
	}

	return true, "Ok!"

}
