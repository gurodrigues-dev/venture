package main

import (
	"gin/routes"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	routes.HandleRequests()

}
