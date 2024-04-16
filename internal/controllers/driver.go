package controllers

import (
	"fmt"
	"gin/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateDriver(c *gin.Context) {

	var input types.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	qrCode, err := ct.service.CreateAndSaveQrCodeInS3(c, &input.CNH)

	if err != nil {
		log.Printf("error to save qrcode: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error whiling creating qrcode"})
		return
	}

	input.QrCode = qrCode

	err = ct.service.CreateDriver(c, &input)

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

	err = ct.service.AddMessageInQueue(c, msg)
	if err != nil {
		log.Printf("error while adding message on queue")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error to send queue"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (ct *controller) ReadDriver(c *gin.Context) {

	cnh := c.Param("cnh")

	driver, err := ct.service.ReadDriver(c, &cnh)

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

	cnhInterface, err := ct.service.ParserJwtDriver(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnh of cookie don't found"})
		return
	}

	cnh, err := ct.service.InterfaceToString(cnhInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the cnh value isn't string"})
		return
	}

	err = ct.service.DeleteDriver(c, cnh)

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

	driver, err := ct.service.AuthDriver(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.service.CreateTokenJWTDriver(c, driver)

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

}

func (ct *controller) CurrentStudents(c *gin.Context) {

}
