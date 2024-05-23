package controllers

import (
	"fmt"
	"gin/types"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateDriver(c *gin.Context) {

	var input types.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	qrCode, err := ct.awsservice.CreateAndSaveQrCodeInS3(c, &input.CNH)

	if err != nil {
		log.Printf("error to save qrcode: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error whiling creating qrcode"})
		return
	}

	input.QrCode = qrCode

	err = ct.driverservice.CreateDriver(c, &input)

	if err != nil {
		log.Printf("error to create driver: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating driver"})
		return
	}

	email := types.Email{
		Recipient: input.Email,
		Subject:   fmt.Sprintf("Verification Email - %s", input.Name),
		Body:      fmt.Sprintf("Hello, my new driver %s! This email was registred in Venture. Our Apprechiated your choose, the choose of technology", input.Name),
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

func (ct *controller) ReadDriver(c *gin.Context) {

	cnh := c.Param("cnh")

	driver, err := ct.driverservice.ReadDriver(c, &cnh)

	if err != nil {
		log.Printf("error while found driver: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "driver don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"driver": driver})

}

func (ct *controller) UpdateDriver(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *controller) DeleteDriver(c *gin.Context) {

	cnhInterface, err := ct.driverservice.ParserJwtDriver(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnh of cookie don't found"})
		return
	}

	cnh, err := ct.service.InterfaceToString(cnhInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the cnh value isn't string"})
		return
	}

	err = ct.driverservice.DeleteDriver(c, cnh)

	if err != nil {
		log.Printf("error whiling deleted driver: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted driver"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, gin.H{"message": "driver deleted w successfully"})

}

func (ct *controller) AuthDriver(c *gin.Context) {

	var input types.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	driver, err := ct.driverservice.AuthDriver(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.driverservice.CreateTokenJWTDriver(c, driver)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"driver": driver,
		"token":  jwt,
	})

}

func (ct *controller) CurrentWorkplaces(c *gin.Context) {

	cnhInterface, err := ct.driverservice.ParserJwtDriver(c)

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

	schools, err := ct.driverservice.GetWorkplaces(c, cnh)
	if err != nil {
		log.Printf("error to search schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error at find schools"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schools": schools})

}

func (ct *controller) CurrentStudents(c *gin.Context) {

}

func (ct *controller) ReadAllInvites(c *gin.Context) {

	cnhInterface, err := ct.driverservice.ParserJwtDriver(c)

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

	invites, err := ct.driverservice.ReadAllInvites(c, cnh)

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

	err = ct.driverservice.UpdateInvite(c, &id)
	if err != nil {
		log.Printf("error while updating invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at updating invite"})
		return
	}

	invite, err := ct.driverservice.ReadInvite(c, &id)
	if err != nil {
		log.Printf("error while reading invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at reading invite"})
		return
	}

	err = ct.driverservice.CreateEmployee(c, invite)
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

	err = ct.kafkaservice.AddMessageInQueue(c, msgSchool)
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

	err = ct.kafkaservice.AddMessageInQueue(c, msgDriver)
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

	err = ct.driverservice.DeleteInvite(c, &id)
	if err != nil {
		log.Printf("error while deleting invite: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error at deleting invite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invite was deleted w/ sucessfully"})

}
