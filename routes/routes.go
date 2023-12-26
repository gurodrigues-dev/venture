package routes

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleRequests() {

	r := gin.Default()

	r.Use(func(c *gin.Context) {

		requestID := uuid.New()

		c.Writer.Header().Set("X-Request-ID", requestID.String())

		c.Set("RequestID", requestID)

		c.Next()

	})

	// health

	r.GET("api/v1/health", controllers.Health)

	// usuarios

	r.POST("api/v1/users", controllers.CreateUser)

	r.GET("api/v1/users/:cpf", controllers.GetUser)

	r.PUT("api/v1/users/:cpf", controllers.UpdateUser)

	r.DELETE("api/v1/users/:cpf", controllers.DeleteUser)

	r.Run()

}
