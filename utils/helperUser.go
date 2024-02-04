package utils

import (
	"fmt"
	"gin/logs"
	"gin/models"
	"gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserAndAdressFromRequest(c *gin.Context) (*models.CreateUser, error) {

	hashedPassword := HashPassword(c.PostForm("password"))

	user := &models.CreateUser{
		Name:     c.PostForm("name"),
		Password: hashedPassword,
		CPF:      c.PostForm("cpf"),
		RG:       c.PostForm("rg"),
		ID:       uuid.New(),
		Email:    c.PostForm("email"),
		Endereco: models.Endereco{
			Rua:         c.PostForm("rua"),
			Numero:      c.PostForm("numero"),
			Complemento: c.PostForm("complemento"),
			Cidade:      c.PostForm("cidade"),
			Estado:      c.PostForm("estado"),
			CEP:         c.PostForm("cep"),
		},
	}

	if user.Endereco.Estado != "sp" && user.Endereco.Estado != "SP" {

		return &models.CreateUser{}, fmt.Errorf("Fora de disponibilidade")

	}

	return user, nil

}

func VerifyUserAndPassword(c *gin.Context) (bool, error) {

	cpf := c.PostForm("cpf")
	table := c.PostForm("table")

	hashedPassword := HashPassword(c.PostForm("password"))

	match, err := repository.VerifyPasswordByCpf(cpf, table, hashedPassword)

	return match, err

}

func VerifyCpf(c *gin.Context) (bool, error) {

	requestData := logs.GetDataOfRequest(c)

	cpfJwtToken, found := c.Get("cpf")

	if !found {
		return false, fmt.Errorf("Erro ao encontrar Token")
	}

	cpfRequest := c.Param("cpf")

	cpfMatch := cpfJwtToken == cpfRequest

	requestData.CPFRequest = cpfRequest
	requestData.CPFJwtToken = cpfJwtToken
	requestData.CPFMatch = cpfMatch

	_, err := logs.LoggingDataOfRequest(requestData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error loading variables",
			"error":  err.Error(),
		})

		return false, fmt.Errorf("Erro ao logar requisição")
	}

	return cpfMatch, nil
}

func ValidateDocsUser(user *models.CreateUser) (bool, string) {

	validateCPF := IsCPF(user.CPF)

	if !validateCPF {

		return false, "cpf invalid, type and try again."

	}

	validateCEP := IsCEP(user.Endereco.CEP)

	if !validateCEP {

		return false, "cep invalid, type and try again."
	}

	return true, "Ok!"

}
