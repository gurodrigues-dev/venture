package controllers

import (
	"fmt"
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

	var input types.ValidaInvite

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err = ct.service.IsEmployee(c, &input.Guest)

	if err != nil {
		log.Printf("already connection: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "you already have a connection with this driver."})
		return
	}

	driver, err := ct.service.ReadDriver(c, &input.Guest)
	if err != nil {
		log.Printf("already connection: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "you already have a connection with this driver."})
		return
	}

	school, err := ct.service.ReadSchool(c, cnpj)
	if err != nil {
		log.Printf("already connection: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "you already have a connection with this driver."})
		return
	}

	invite := types.Invite{
		School: *school,
		Driver: *driver,
	}

	err = ct.service.CreateInvite(c, &invite)

	if err != nil {
		log.Printf("error to create invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at creating invite"})
		return
	}

	emailSchool := types.Email{
		Recipient: invite.School.Email,
		Subject:   fmt.Sprintf("Invite sended to - %s", invite.Driver.Name),
		Body:      fmt.Sprintf("Hello, %s! you just sent an invite to %s", invite.School.Email, invite.Driver.Name),
	}

	msgSchool, err := ct.service.EmailStructToJSON(&emailSchool)
	if err != nil {
		log.Printf("error while convert email to message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to parse json email"})
		return
	}

	log.Print("mensagem enviada para fila -> ", msgSchool)

	err = ct.service.AddMessageInQueue(c, msgSchool)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	emailDriver := types.Email{
		Recipient: invite.Driver.Email,
		Subject:   fmt.Sprintf("You received an invite of %s", invite.School.Name),
		Body:      fmt.Sprintf("Hello, %s! This email was showing, who sended a invite for you, verify your invites on platform", invite.Driver.Name),
	}

	msgDriver, err := ct.service.EmailStructToJSON(&emailDriver)
	if err != nil {
		log.Printf("error while convert email to message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to parse json email"})
		return
	}

	err = ct.service.AddMessageInQueue(c, msgDriver)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	log.Print("mensagem enviada para fila -> ", msgDriver)

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

	err = ct.service.CreateEmployee(c, invite)
	if err != nil {
		log.Printf("error while creating record of bond: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at creating record of bond"})
		return
	}

	emailSchool := types.Email{
		Recipient: invite.School.Email,
		Subject:   fmt.Sprintf("Invite accepted by - %s", invite.Driver.Name),
		Body:      fmt.Sprintf("Hello, %s! your invite was accepted, happy employee! %s", invite.School.Email, invite.Driver.Name),
	}

	msgSchool, err := ct.service.EmailStructToJSON(&emailSchool)
	if err != nil {
		log.Printf("error while convert email to message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to parse json email"})
		return
	}

	log.Print("mensagem enviada para fila -> ", msgSchool)

	err = ct.service.AddMessageInQueue(c, msgSchool)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	emailDriver := types.Email{
		Recipient: invite.Driver.Email,
		Subject:   fmt.Sprintf("You accepted invite of %s", invite.School.Name),
		Body:      fmt.Sprintf("Hello, %s! Congratulations, you created at new partner and a new workplace, cheers!", invite.Driver.Name),
	}

	msgDriver, err := ct.service.EmailStructToJSON(&emailDriver)
	if err != nil {
		log.Printf("error while convert email to message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to parse json email"})
		return
	}

	err = ct.service.AddMessageInQueue(c, msgDriver)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	log.Print("mensagem enviada para fila -> ", msgDriver)

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
