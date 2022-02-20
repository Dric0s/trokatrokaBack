package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"trokatrokaBack/controller"
)

func main() {
	router := gin.Default()
	controller.InitializateRoutes(router)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Wasn't possible to run the server in default port: 8080")
	}
}
