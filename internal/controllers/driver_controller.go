package controllers

import (
	"gin/config"
	"gin/internal/service"
	"gin/types"
	"gin/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ClaimsDriver struct {
	CNH string `json:"cnh"`
	jwt.StandardClaims
}

type DriverController struct {
	driverservice *service.DriverService
}

func NewDriverController(service *service.DriverService) *DriverController {
	return &DriverController{driverservice: service}
}

func (ct *DriverController) RegisterRoutes(router *gin.Engine) {

	conf := config.Get()

	driverMiddleware := func(c *gin.Context) {

		secret := []byte(conf.Server.Secret)

		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sem cookie de sessão"})
			c.Abort()
			return
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &ClaimsDriver{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*ClaimsDriver)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("cnh", claims.CNH)
		c.Set("isAuthenticated", true)
		c.Next()

	}

	api := router.Group("api/v1")

	api.POST("/driver", ct.CreateDriver)
	api.GET("/driver/:cnh", ct.ReadDriver)
	api.PATCH("/driver", driverMiddleware, ct.UpdateDriver)
	api.DELETE("/driver", driverMiddleware, ct.DeleteDriver)
	api.POST("/login/driver", ct.AuthDriver)
	api.GET("/driver/partners", driverMiddleware, ct.CurrentWorkplaces)
	api.GET("/driver/sponsors", driverMiddleware) // yet doesnt deployed
	api.GET("/invite", driverMiddleware, ct.ReadAllInvites)
	api.PATCH("/invite/:id", driverMiddleware, ct.UpdateInvite)
	api.DELETE("/invite/:id", driverMiddleware, ct.DeleteInvite)

}

func (ct *DriverController) CreateDriver(c *gin.Context) {

	var input types.Driver

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	qrCode, err := ct.driverservice.CreateAndSaveQrCodeInS3(c, &input.CNH)

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

	c.JSON(http.StatusCreated, input)
}

func (ct *DriverController) ReadDriver(c *gin.Context) {

	cnh := c.Param("cnh")

	driver, err := ct.driverservice.ReadDriver(c, &cnh)

	if err != nil {
		log.Printf("error while found driver: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "driver don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"driver": driver})

}

func (ct *DriverController) UpdateDriver(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *DriverController) DeleteDriver(c *gin.Context) {

	cnhInterface, err := ct.driverservice.ParserJwtDriver(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnh of cookie don't found"})
		return
	}

	cnh, err := utils.InterfaceToString(cnhInterface)

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

func (ct *DriverController) AuthDriver(c *gin.Context) {

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

func (ct *DriverController) CurrentWorkplaces(c *gin.Context) {

	cnhInterface, err := ct.driverservice.ParserJwtDriver(c)

	if err != nil {
		log.Printf("error to read jwt: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of jwt"})
		return
	}

	cnh, err := utils.InterfaceToString(cnhInterface)

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

func (ct *DriverController) CurrentStudents(c *gin.Context) {

}

func (ct *DriverController) ReadAllInvites(c *gin.Context) {

	cnhInterface, err := ct.driverservice.ParserJwtDriver(c)

	if err != nil {
		log.Printf("error to read jwt: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of jwt"})
		return
	}

	cnh, err := utils.InterfaceToString(cnhInterface)

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

func (ct *DriverController) UpdateInvite(c *gin.Context) {

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

	c.JSON(http.StatusCreated, invite)

}

func (ct *DriverController) DeleteInvite(c *gin.Context) {

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
