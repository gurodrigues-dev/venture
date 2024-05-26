package controllers

import (
	"gin/config"
	"gin/internal/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ChildController struct {
	childservice *service.ChildService
}

func NewChildController(service *service.ChildService) *ChildController {
	return &ChildController{childservice: service}
}

func (ct *ChildController) RegisterRoutes(router *gin.Engine) {

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

	api.POST("/child", authMiddleware, ct.CreateChildren)
	api.GET("/child", authMiddleware, ct.ReadChildren)
	api.PATCH("/child", authMiddleware, ct.UpdateChildren)
	api.DELETE("/child", authMiddleware, ct.DeleteChildren)

}

func (ct *ChildController) CreateChildren(c *gin.Context) {

}

func (ct *ChildController) ReadChildren(c *gin.Context) {

}

func (ct *ChildController) UpdateChildren(c *gin.Context) {

}

func (ct *ChildController) DeleteChildren(c *gin.Context) {

}
