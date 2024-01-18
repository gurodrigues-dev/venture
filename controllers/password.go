package controllers

import (
	"fmt"
	"gin/models"
	"gin/repository"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryPassword(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	email := c.PostForm("email")

	emailExistsUsers, err := repository.CheckExistsEmailInUsers(email)

	emailExistsDrivers, err := repository.CheckExistsEmailInDrivers(email)

	if !emailExistsUsers && !emailExistsDrivers {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email not found.",
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
		"message":   "token generated successfully",
		"redis-log": "key and value received",
		"email-log": "email sended success",
		"requestid": requestID,
	})

	return

}

func VerifyIdentityToChangePassword(c *gin.Context) {

	requestID, _ := c.Get("requestID")

	user := models.UserInfoToResetPassword{
		Token: c.PostForm("token"),
		Email: c.PostForm("email"),
	}

	validate, err := repository.VerifyMatchTokensToResetPassword(user.Email, user.Token)

	if !validate {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":   err.Error(),
			"requestID": requestID,
			"status":    "Token incorreto.",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "redis authenticated token ",
		"requestid": requestID,
	})

	return

}

func ChangePassword(c *gin.Context) {

	requestID, _ := c.Get("requestID")

	hashedPassword := utils.HashPassword(c.PostForm("password"))

	user := models.UserResetPassword{
		Email:           c.PostForm("email"),
		NewHashPassword: hashedPassword,
	}

	resp, err := repository.ChangePasswordByEmailIdentification(user)

	if !resp {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Erro ao salvar nova senha no banco",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "password updated w/ sucess",
		"requestid": requestID,
	})

	return

}
