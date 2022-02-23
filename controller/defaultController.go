package controller

import "github.com/gin-gonic/gin"

func InitializateRoutes(router *gin.Engine) {
	//Book
	router.POST("/book", PostBook)
	router.POST("/login", Login)
}
