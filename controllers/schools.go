package controllers

import (
	"gin/models"
	"gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSchool(c *gin.Context) {

	var school models.School

	err := c.ShouldBindJSON(&school)

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
