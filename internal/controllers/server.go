package controllers

import (
	"fmt"
	"gin/config"
	"gin/internal/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type controller struct {
	service *service.Service
}

func New(s *service.Service) *controller {
	return &controller{
		service: s,
	}
}

type ClaimsUser struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}

type ClaimsDriver struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}

type ClaimsSchool struct {
	CNPJ string `json:"cnpj"`
	jwt.StandardClaims
}

// @Summary	Show API ping
//
//	@Success	200				{string}	string
//
// @Router		/api/v1/ping [get]
func (ct *controller) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (ct *controller) Start() {

	conf := config.Get()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// middleware for users
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

		token, err := jwt.ParseWithClaims(tokenString, &ClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*ClaimsUser)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("cpf", claims.CPF)
		c.Set("isAuthenticated", true)
		c.Next()

	}

	schoolMiddleware := func(c *gin.Context) {

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

		token, err := jwt.ParseWithClaims(tokenString, &ClaimsSchool{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*ClaimsSchool)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("cnpj", claims.CNPJ)
		c.Set("isAuthenticated", true)
		c.Next()

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1")
	api.GET("/ping", ct.ping)

	// user
	api.POST("/user")
	api.GET("/user", authMiddleware)
	api.PATCH("/user", authMiddleware)
	api.DELETE("/user", authMiddleware)

	// child
	api.POST("/child", authMiddleware)
	api.GET("/child", authMiddleware)
	api.PATCH("/child", authMiddleware)
	api.DELETE("/child", authMiddleware)

	// driver
	api.POST("/driver", authMiddleware)
	api.GET("/driver", authMiddleware)
	api.PATCH("/driver", authMiddleware)
	api.DELETE("/driver", authMiddleware)

	// school
	api.POST("/school")
	api.GET("/school", schoolMiddleware)
	api.PATCH("/school", schoolMiddleware)
	api.DELETE("/school", schoolMiddleware)

	router.Run(fmt.Sprintf(":%d", conf.Server.Port))
}
