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

	api.POST("/responsible", ct.CreateResponsible)                   // criar conta
	api.GET("/responsible/:cpf", ct.ReadResponsible)                 // entrar no perfil
	api.PATCH("/responsible", authMiddleware, ct.UpdateResponsible)  // atualizar conta
	api.DELETE("/responsible", authMiddleware, ct.DeleteResponsible) // deletar minha conta
	api.POST("/login/responsible", ct.AuthResponsible)               // login de responsavel
	api.POST("/child", authMiddleware, ct.CreateChild)               // registrar filho
	api.GET("/child", authMiddleware, ct.ReadChildren)               // verificar todos os filhos
	api.PATCH("/child", authMiddleware, ct.UpdateChild)              // atualizar infos sobre o filho
	api.DELETE("/child/:rg", authMiddleware, ct.DeleteChild)         // deletar um filho
	api.GET("/:rg/schools", authMiddleware)                          // verificar todas as escolas pára registrar meu filho
	api.GET("/:rg/:cnpj/drivers", authMiddleware)                    // verificar todos os motoristas da escola para registrar meu filho
	api.GET("/:rg/:cnpj/:cnh", authMiddleware)                       // verificando infos sobre um motorista pra continuar o registro
	api.POST("/:rg/:cnpj/:cnh", authMiddleware)                      // criando uma matricula na escola e assinando contrato com motorista
	api.DELETE("/sponsor/:id", authMiddleware)                       // deletar vinculo com motorista e escola

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
		log.Printf("error whiling deleted responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted responsible"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, gin.H{"message": "responsible deleted w successfully"})
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

func (ct *ResponsibleController) CreateChild(c *gin.Context) {

	var child types.Child
	if err := c.ShouldBindJSON(&child); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	child.Responsible.CPF = *cpf

	err = ct.responsibleservice.CreateChild(c, &child)

	if err != nil {
		log.Printf("error to create child: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error at create child"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"child": &child})

}

func (ct *ResponsibleController) ReadChildren(c *gin.Context) {

	cpfInterface, err := ct.responsibleservice.ParserJwtResponsible(c)

	if err != nil {
		log.Printf("error to parse jwt: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "cpf of cookie don't found"})
		return
	}

	cpf, err := ct.responsibleservice.InterfaceToString(cpfInterface)

	if err != nil {
		log.Printf("error to parse interface: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "the cpf value isn't string"})
		return
	}

	children, err := ct.responsibleservice.ReadChildren(c, cpf)

	if err != nil {
		log.Printf("error to search children: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "children don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"children": children})

}

func (ct *ResponsibleController) UpdateChild(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "updated w success"})

}

func (ct *ResponsibleController) DeleteChild(c *gin.Context) {

	rg := c.Param("rg")

	sponsor := ct.responsibleservice.IsSponsor(c, &rg)

	if sponsor {
		log.Print("this child is sponsor")
		c.JSON(http.StatusNotModified, gin.H{"error": "error at deleting, this child is sponsor"})
		return
	}

	err := ct.responsibleservice.DeleteChild(c, &rg)

	if err != nil {
		log.Printf("error while deleting responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "child doesnt deleting"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleting w success"})

}

func (ct *ResponsibleController) GetSchools(c *gin.Context) {

}

func (ct *ResponsibleController) GetDriversInSchools(c *gin.Context) {

}

func (ct *ResponsibleController) GetDriver(c *gin.Context) {

}

func (ct *ResponsibleController) CreateSponsor(c *gin.Context) {

}

func (ct *ResponsibleController) DeleteSponsor(c *gin.Context) {

}
