package controllers

import (
	"gin/models"
	"gin/repository"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSchool(c *gin.Context) {

	var school models.School

	err := c.ShouldBindJSON(&school)

	emailExist, err := repository.CheckExistsEmailInUsers(school.Email)

	if emailExist {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Este email já existe.",
			"error":   "email found.",
		})

		return
	}

	_, err = utils.SendMessageOfVerifyToEmailInAwsSes(school.Email)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao encontrar email.",
			"error":   err.Error(),
		})

		return

	}

	school.Password = utils.HashPassword(school.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = repository.SaveSchool(&school)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"data":    &school,
		"message": "school created w/ success",
	})

	return
}

func GetSchool(c *gin.Context) {

	resp, _ := utils.VerifyCpf(c)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Security breach, intruder account trying to delete account.",
			"message": "Invalid Cnpj",
		})

		return

	}

	return

}

func UpdateSchool(c *gin.Context) {

	return

}

func DeleteSchool(c *gin.Context) {

	resp, _ := utils.VerifyCnpj(c)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Security breach, intruder account trying to delete account.",
			"message": "Invalid Cnpj",
		})

		return

	}

	cnpj := c.Param("cnpj")

	emailToDelete, err := repository.DeleteSchoolByCnpj(&cnpj)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error while deleting in database",
		})

		return
	}

	_, err = utils.DeleteEmailFromAwsSes(&emailToDelete)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Error when deleting user email of SES",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted w/ success",
	})

	return

}

func AuthenticateSchool(c *gin.Context) {

	var LoginInfo models.LoginSchool

	err := c.ShouldBindJSON(&LoginInfo)

	_, err = utils.VerifySchoolAndPassword(&LoginInfo)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Login error",
			"error":   err.Error(),
		})

		return
	}

	tokenJwt, err := utils.CreateJwtToken(LoginInfo.CNPJ)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating JWToken",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "login accepted",
		"token":   tokenJwt,
	})

	return

}
