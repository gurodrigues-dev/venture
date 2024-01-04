package controllers

import (
	"gin/repository"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResetPassword(c *gin.Context) {

	return

}

func RecoveryPassword(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	email := c.PostForm("email")

	emailExists, err := repository.CheckExistsEmail(email)

	if !emailExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email not exist found.",
			"error":   err.Error(),
		})

		return
	}

	token, err := utils.GenerateRandomToken()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	resp, err := repository.SaveTokenToRedis(email, token)

	if !resp {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar token no Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Token gerado com sucesso",
		"redis-log": "Key and value received",
		"email-log": "Email sended success",
		"requestid": requestID,
	})

	return

}
