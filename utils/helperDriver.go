package utils

import (
	"gin/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDriverAndAdressFromRequest(c *gin.Context, url string) (*models.CreateDriver, *models.Endereco) {

	hashedPassword := HashPassword(c.PostForm("password"))

	driver := &models.CreateDriver{
		Name:     c.PostForm("name"),
		Password: hashedPassword,
		CPF:      c.PostForm("cpf"),
		RG:       c.PostForm("rg"),
		CNH:      c.PostForm("cnh"),
		ID:       uuid.New(),
		Email:    c.PostForm("email"),
		URL:      url,
	}

	endereco := &models.Endereco{
		Rua:         c.PostForm("rua"),
		Numero:      c.PostForm("numero"),
		Complemento: c.PostForm("complemento"),
		Cidade:      c.PostForm("cidade"),
		Estado:      c.PostForm("estado"),
		CEP:         c.PostForm("cep"),
	}

	return driver, endereco

}

func ValidateDocsDriver(user *models.CreateDriver, endereco *models.Endereco) (bool, string) {

	validateCPF := IsCPF(user.CPF)

	if !validateCPF {

		return false, "cpf invalid, type and try again."

	}

	validateCNH := IsCNH(user.CNH)

	if !validateCNH {

		return false, "cnh invalid, type and try again."
	}

	validateCEP := IsCEP(endereco.CEP)

	if !validateCEP {

		return false, "cep invalid, type and try again."
	}

	return true, "Ok!"

}
