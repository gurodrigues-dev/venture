package controllers

import (
	"fmt"
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

	resp, err := repository.SaveClient(user, endereco)

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

	user, err := repository.FindByCpf(cpf)

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

	requestID, _ := c.Get("RequestID")

	resp, ok := utils.VerifyCpf(c)

	fmt.Println(resp, ok)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     "Security breach, intruder account trying to delete account.",
			"message":   "Invalid Cpf",
		})

		return

	}

}

func DeleteUser(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	resp, ok := utils.VerifyCpf(c)

	fmt.Println(resp, ok)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     "Security breach, intruder account trying to delete account.",
			"message":   "Invalid Cpf",
		})

		return

	}

	cpf := c.Param("cpf")

	_, err := repository.DeleteByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"requestID": requestID,
			"error":     err.Error(),
			"message":   "Error while deleting in database",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requestID": requestID,
		"message":   "User deleted w/ success",
	})

}

func AuthenticateUser(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	resp, err := utils.VerifyUserAndPassword(c)

	if !resp {
		c.JSON(http.StatusUnauthorized, gin.H{
			"requestID": requestID,
			"message":   "Login error",
			"error":     err.Error(),
		})

		return
	}

	tokenJwt, err := utils.CreateJwtToken(c.PostForm("cpf"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"requestID": requestID,
			"message":   "Error while creating JWToken",
			"error":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":   "login accepted",
		"requestID": requestID,
		"token":     tokenJwt,
	})

}
