package controllers

import (
	"gin/config"
	"gin/internal/service"
	"gin/types"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ClaimsResponsible struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}

type ResponsibleController struct {
	responsibleservice *service.ResponsibleService
}

func NewResponsibleController(service *service.ResponsibleService) *ResponsibleController {
	return &ResponsibleController{responsibleservice: service}
}

func (ct *ResponsibleController) RegisterRoutes(router *gin.Engine) {

	conf := config.Get()

	authMiddleware := func(c *gin.Context) {

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

		token, err := jwt.ParseWithClaims(tokenString, &ClaimsResponsible{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*ClaimsResponsible)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("cpf", claims.CPF)
		c.Set("isAuthenticated", true)
		c.Next()

	}

	api := router.Group("api/v1")

	api.POST("/responsible", ct.CreateResponsible)
	api.GET("/responsible", ct.ReadResponsible)
	api.PATCH("/responsible", authMiddleware, ct.UpdateResponsible)
	api.DELETE("/responsible", authMiddleware, ct.DeleteResponsible)

}

func (ct *ResponsibleController) CreateResponsible(c *gin.Context) {
	var responsible types.Responsible
	if err := c.ShouldBindJSON(&responsible); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ct.responsibleservice.CreateResponsible(c, &responsible)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, responsible)
}

func (ct *ResponsibleController) ReadResponsible(c *gin.Context) {

	cpf := c.Param("cpf")

	responsible, err := ct.responsibleservice.ReadResponsible(c, &cpf)

	if err != nil {
		log.Printf("error while found responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "responsible don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"responsible": responsible})

}

func (ct *ResponsibleController) UpdateResponsible(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})
}

func (ct *ResponsibleController) DeleteResponsible(c *gin.Context) {
	cpfInterface, err := ct.responsibleservice.ParserJwtResponsible(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cpf of cookie don't found"})
		return
	}

	cpf, err := ct.responsibleservice.InterfaceToString(cpfInterface)

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

func (ct *ResponsibleController) AuthResponsible(c *gin.Context) {
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
