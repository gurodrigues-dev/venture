package controllers

import (
	"gin/repository"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	respOfAws, err := utils.SaveQRCodeOfUser(c.PostForm("cpf"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": respOfAws,
			"error":   err.Error(),
		})

		return
	}

	user, endereco := utils.GetUserAndAdressFromRequest(c, respOfAws)

	validateDocs, documentError := utils.ValidateDocsDriver(user, endereco)

	if !validateDocs {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "type and try insert your documents again, please.",
			"error":   documentError,
		})

		return

	}

	resp, err := repository.InsertNewUser(user, endereco)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error when inserting into database",
			"error":   err.Error(),
		})

		return

	}

	if resp == true {

		c.JSON(http.StatusCreated, gin.H{
			"requestID":   requestID,
			"status":      "user created successfully",
			"s3bucketurl": respOfAws,
		})
	}

	return

}

func GetUser(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	cpf := c.Param("cpf")

	user, err := repository.GetUser(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error while searching in database",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requestID": requestID,
		"userData":  user,
	})

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
