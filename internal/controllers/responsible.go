package controllers

import (
	"gin/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateResponsible(c *gin.Context) {

	var input types.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.responsibleservice.CreateResponsible(c, &input)

	if err != nil {
		log.Printf("error to create driver: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating driver"})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *controller) ReadResponsible(c *gin.Context) {

	cpf := c.Param("cpf")

	responsible, err := ct.responsibleservice.ReadResponsible(c, &cpf)

	if err != nil {
		log.Printf("error while found responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "responsible don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"responsible": responsible})

}

func (ct *controller) UpdateResponsible(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})
}

func (ct *controller) DeleteResponsible(c *gin.Context) {
	cpfInterface, err := ct.responsibleservice.ParserJwtResponsible(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cpf of cookie don't found"})
		return
	}

	cpf, err := ct.service.InterfaceToString(cpfInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the cpf value isn't string"})
		return
	}

	err = ct.responsibleservice.DeleteResponsible(c, cpf)

	if err != nil {
		log.Printf("error whiling deleted driver: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted driver"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, gin.H{"message": "driver deleted w successfully"})
}

func (ct *controller) AuthResponsible(c *gin.Context) {
	var input types.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	responsible, err := ct.responsibleservice.AuthResponsible(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.responsibleservice.CreateTokenJWTResponsible(c, responsible)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"responsible": responsible,
		"token":       jwt,
	})
}
