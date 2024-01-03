package controllers

import (
	"gin/config"
	"gin/repository"
	"gin/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func ResetPassword(c *gin.Context) {

	return

}

func RecoveryPassword(c *gin.Context) {

	config.LoadEnvironmentVariables()

	var (
		redisAddress  = config.GetRedisAddress()
		redisPassword = config.GetRedisPassword()
	)

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

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})

	err = client.Set(email, token, 10*time.Minute).Err()

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar token no Redis"})
		return
	}

	// enviar email com aws SES

	return

}
