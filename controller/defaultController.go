package controller

import "github.com/gin-gonic/gin"

func InitializateRoutes(router *gin.Engine) {
	//Book
	router.POST("/book", postBook)
}
