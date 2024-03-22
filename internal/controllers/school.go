package controllers

import (
	"gin/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("Erro ao parsear o body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Conteúdo do body inválido"})
	}

	err := ct.service.CreateSchool(c, &input)

	if err != nil {
		log.Printf("Erro ao criar escola: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Ocorreu um erro ao criar a conta."})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *controller) ReadSchool(c *gin.Context) {

}

func (ct *controller) UpdateSchool(c *gin.Context) {

}

func (ct *controller) DeleteSchool(c *gin.Context) {

}

func (ct *controller) AuthSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("Erro ao parsear o body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Conteúdo do body inválido"})
		return
	}

	school, err := ct.service.AuthSchool(c, &input)

	if err != nil {
		log.Printf("Email ou senha incorreto: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Email ou senha incorreto."})
		return
	}

	jwt, err := ct.service.CreateTokenJWTSchool(c, &input)

	if err != nil {
		log.Printf("Erro ao criar token JWT: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro ao criar Token de Autenticação"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})

}
