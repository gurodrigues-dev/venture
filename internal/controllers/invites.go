package controllers

import (
	"gin/types"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateInvite(c *gin.Context) {

	cnpjInterface, err := ct.service.ParserJwtSchool(c)

	if err != nil {
		log.Printf("error to read jwt: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of jwt"})
		return
	}

	cnpj, err := ct.service.InterfaceToString(cnpjInterface)

	if err != nil {
		log.Printf("error to convert interface in string: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	var input types.Invite

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	// verificar se eles ja nao tem vinculo para nao duplicar invites

	input.Requester = *cnpj

	err = ct.service.CreateInvite(c, &input)

	if err != nil {
		log.Printf("error to create invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at creating invite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invite sended was successfully"})

}

// The middleware has different of CreateInvite
func (ct *controller) ReadAllInvites(c *gin.Context) {

	cnhInterface, err := ct.service.ParserJwtDriver(c)

	if err != nil {
		log.Printf("error to read jwt: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of jwt"})
		return
	}

	cnh, err := ct.service.InterfaceToString(cnhInterface)

	if err != nil {
		log.Printf("error to convert interface in string: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	invites, err := ct.service.ReadAllInvites(c, cnh)

	if err != nil {
		log.Printf("error while found invites: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invites don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"invites": invites})

}

func (ct *controller) UpdateInvite(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("error while convert string to int: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at convert to int"})
		return
	}

	err = ct.service.UpdateInvite(c, &id)
	if err != nil {
		log.Printf("error while updating invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at updating invite"})
		return
	}

	invite, err := ct.service.ReadInvite(c, &id)
	if err != nil {
		log.Printf("error while reading invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at reading invite"})
		return
	}

	err = ct.service.CreateRecordToSchoolAndDriver(c, invite)
	if err != nil {
		log.Printf("error while creating record of bond: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at creating record of bond"})
		return
	}

	c.JSON(http.StatusCreated, invite)

}

func (ct *controller) DeleteInvite(c *gin.Context) {

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("error while convert string to int: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at convert to int"})
		return
	}

	err = ct.service.DeleteInvite(c, &id)
	if err != nil {
		log.Printf("error while deleting invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at deleting invite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invite was deleted w/ sucessfully"})

}
