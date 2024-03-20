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
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
	}

}

func (ct *controller) ReadSchool(c *gin.Context) {

}

func (ct *controller) UpdateSchool(c *gin.Context) {

}

func (ct *controller) DeleteSchool(c *gin.Context) {

}

func (ct *controller) AuthSchool(c *gin.Context) {

}
