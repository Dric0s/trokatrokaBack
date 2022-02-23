package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trokatrokaBack/model"
)

func PostBook(c *gin.Context) {
	var newBook model.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}
