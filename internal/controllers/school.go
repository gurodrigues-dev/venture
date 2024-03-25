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
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.service.CreateSchool(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating school"})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *controller) ReadSchool(c *gin.Context) {

	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusNotFound, "cookie don't found")
		return
	}

	cnpj, err := ct.service.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, "name don't found")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"teste":  "cookie validated",
		"cookie": cookie,
		"cnpj":   cnpj,
	})

}

func (ct *controller) UpdateSchool(c *gin.Context) {

}

func (ct *controller) DeleteSchool(c *gin.Context) {

}

func (ct *controller) AuthSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	school, err := ct.service.AuthSchool(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.service.CreateTokenJWTSchool(c, school)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})

}
