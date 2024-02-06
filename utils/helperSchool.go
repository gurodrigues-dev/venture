package utils

import (
	"fmt"
	"gin/models"
	"gin/repository"

	"github.com/gin-gonic/gin"
)

func VerifySchoolAndPassword(data *models.LoginSchool) (bool, error) {

	hashedPassword := HashPassword(data.Password)

	match, err := repository.VerifyPasswordByCnpj(&data.CNPJ, &data.Type, &hashedPassword)

	return match, err

}

func VerifyCnpj(c *gin.Context) (bool, error) {

	cnpjJwtToken, found := c.Get("cnpj")

	if !found {
		return false, fmt.Errorf("Erro ao encontrar Token")
	}

	cnpjRequest := c.Param("cnpj")

	cnpjMatch := cnpjJwtToken == cnpjRequest

	return cnpjMatch, nil
}
