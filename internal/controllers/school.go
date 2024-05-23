package controllers

import (
	"fmt"
	"gin/types"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.schoolservice.CreateSchool(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating school"})
		return
	}

	email := types.Email{
		Recipient: input.Email,
		Subject:   fmt.Sprintf("Verification Email - %s", input.Name),
		Body:      fmt.Sprintf("Hello, %s! This email was registred in Venture. Our Apprechiated your choose, the choose of technology", input.Name),
	}

	msg, err := ct.service.EmailStructToJSON(&email)
	if err != nil {
		log.Printf("error while convert email to message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to parse json email"})
		return
	}

	log.Print("mensagem enviada para fila -> ", msg)

	err = ct.kafkaservice.AddMessageInQueue(c, msg)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *controller) ReadSchool(c *gin.Context) {

	cnpj := c.Param("cnpj")

	school, err := ct.schoolservice.ReadSchool(c, &cnpj)

	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"school": school})

}

func (ct *controller) ReadAllSchools(c *gin.Context) {

	cnpj, err := ct.schoolservice.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, "cnpj of cookie don't found")
		return
	}

	log.Print("consulting page -->", cnpj)

	schools, err := ct.schoolservice.ReadAllSchools(c)

	if err != nil {
		log.Printf("error while found schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "schools don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schools": schools})

}

func (ct *controller) UpdateSchool(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *controller) DeleteSchool(c *gin.Context) {

	cnpjInterface, err := ct.schoolservice.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnpj of cookie don't found"})
		return
	}

	cnpj, err := ct.service.InterfaceToString(cnpjInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	err = ct.schoolservice.DeleteSchool(c, cnpj)

	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, gin.H{"message": "school deleted w successfully"})

}

func (ct *controller) AuthSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	school, err := ct.schoolservice.AuthSchool(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.schoolservice.CreateTokenJWTSchool(c, school)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})
}

func (ct *controller) GetEmployees(c *gin.Context) {

	cnpjInterface, err := ct.schoolservice.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnpj of cookie don't found"})
		return
	}

	cnpj, err := ct.service.InterfaceToString(cnpjInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	drivers, err := ct.schoolservice.GetEmployees(c, cnpj)

	if err != nil {
		log.Printf("error while searching drivers: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error at find drivers"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"drivers": drivers})

}

func (ct *controller) DeleteEmployee(c *gin.Context) {

	record_idStr := c.Param("id")

	record_id, err := strconv.Atoi(record_idStr)

	if err != nil {
		log.Printf("error while converting record_id of string to int: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
		return
	}

	err = ct.schoolservice.DeleteEmployee(c, &record_id)

	if err != nil {
		log.Printf("error while deleting record: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error while deleting employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "employee deleted w/ success"})

}

func (ct *controller) CreateInvite(c *gin.Context) {

	cnpjInterface, err := ct.schoolservice.ParserJwtSchool(c)

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

	err = ct.schoolservice.IsEmployee(c, &input.Guest)

	if err != nil {
		log.Printf("already connection: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "you already have a connection with this driver."})
		return
	}

	driver, err := ct.driverservice.ReadDriver(c, &input.Guest)
	if err != nil {
		log.Printf("already connection: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "you already have a connection with this driver."})
		return
	}

	school, err := ct.schoolservice.ReadSchool(c, cnpj)
	if err != nil {
		log.Printf("already connection: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "you already have a connection with this driver."})
		return
	}

	invite := types.Invite{
		School: *school,
		Driver: *driver,
	}

	err = ct.schoolservice.CreateInvite(c, &invite)

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

	err = ct.kafkaservice.AddMessageInQueue(c, msgSchool)
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

	err = ct.kafkaservice.AddMessageInQueue(c, msgDriver)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	log.Print("mensagem enviada para fila -> ", msgDriver)

	c.JSON(http.StatusOK, gin.H{"message": "invite sended was successfully"})

}
