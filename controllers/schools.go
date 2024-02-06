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

	return

}

func UpdateSchool(c *gin.Context) {

	return

}

func DeleteSchool(c *gin.Context) {

	return

}

func AuthenticateSchool(c *gin.Context) {

	return

}
