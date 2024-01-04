package controllers

import (
	"fmt"
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

	body := fmt.Sprintf("Este é seu token de recuperação: %s, você tem 10 minutos para recuperar sua senha com ele. Caso, o tempo seja ultrapassado, por gentileza. Gere um novo token.", token)

	err = utils.SendEmailAwsSes("Token para recuperação de senha", body, email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   err.Error(),
			"requestid": requestID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Token gerado com sucesso",
		"redis-log": "Key and value received",
		"email-log": "Email sended success",
		"requestid": requestID,
	})

	return

}
