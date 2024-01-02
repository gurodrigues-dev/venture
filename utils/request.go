package utils

import (
	"gin/models"
	"gin/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserAndAdressFromRequest(c *gin.Context, url string) (*models.User, *models.Endereco) {

	hashedPassword := HashPassword(c.PostForm("password"))

	user := &models.User{
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

	return user, endereco

}

func VerifyUserAndPassword(c *gin.Context) (bool, error) {

	cpf := c.PostForm("cpf")

	hashedPassword := HashPassword(c.PostForm("password"))

	match, err := repository.VerifyPasswordByCpf(cpf, hashedPassword)

	return match, err

}
