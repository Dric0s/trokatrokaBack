package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"log"
	"os"
	"trokatrokaBack/controller"
)

var client *redis.Client

func init() {

	dsn := os.Getenv("REDIS_DSN")

	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()
	controller.InitializateRoutes(router)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Wasn't possible to run the server in default port: 8080")
	}
}
