package utils

import (
	"fmt"
	"gin/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDriverAndAdressFromRequest(c *gin.Context, url string) (*models.CreateDriver, error) {

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
		Endereco: models.Endereco{
			Rua:         c.PostForm("rua"),
			Numero:      c.PostForm("numero"),
			Complemento: c.PostForm("complemento"),
			Cidade:      c.PostForm("cidade"),
			Estado:      c.PostForm("estado"),
			CEP:         c.PostForm("cep"),
		},
	}

	if driver.Endereco.Estado != "SP" && driver.Endereco.Estado != "sp" {

		return &models.CreateDriver{}, fmt.Errorf("Fora de disponibilidade")

	}

	return driver, nil

}

func GettingNowInfoFromUserAndRequestInfos(c *gin.Context, user *models.InfoUserToDriver) *models.CreateDriver {

	hashedPassword := HashPassword(c.PostForm("password"))

	driver := &models.CreateDriver{
		Name:     user.Info.Name,
		Password: hashedPassword,
		CPF:      user.Info.CPF,
		RG:       user.Info.RG,
		CNH:      c.PostForm("cnh"),
		ID:       uuid.New(),
		Email:    user.Info.Email,
		URL:      user.URL,
		Endereco: models.Endereco{
			Rua:         user.Info.Endereco.Rua,
			Numero:      user.Info.Endereco.Numero,
			Complemento: user.Info.Endereco.Complemento,
			Cidade:      user.Info.Endereco.Cidade,
			Estado:      user.Info.Endereco.Estado,
			CEP:         user.Info.Endereco.CEP,
		},
	}

	return driver

}

func ValidateDocsDriver(user *models.CreateDriver) (bool, string) {

	validateCPF := IsCPF(user.CPF)

	if !validateCPF {

		return false, "cpf invalid, type and try again."

	}

	validateCNH := IsCNH(user.CNH)

	if !validateCNH {

		return false, "cnh invalid, type and try again."
	}

	validateCEP := IsCEP(user.Endereco.CEP)

	if !validateCEP {

		return false, "cep invalid, type and try again."
	}

	return true, "Ok!"

}
