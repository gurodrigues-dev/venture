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
		return
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

	cookie, err := c.Cookie("username")
	if err != nil {
		c.String(http.StatusNotFound, "Cookie não encontrado")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"teste":  "cookie validado",
		"cookie": cookie,
	})

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
		log.Printf("Email ou senha incorretos: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Email ou senha incorretos."})
		return
	}

	jwt, err := ct.service.CreateTokenJWTSchool(c, &input)

	if err != nil {
		log.Printf("Erro ao criar token JWT: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro ao criar Token de Autenticação"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})

}
