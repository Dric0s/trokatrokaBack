package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trokatrokaBack/model"
	"trokatrokaBack/service"
)

func Login(c *gin.Context) {
	var userAux model.UserData

	if err := c.ShouldBindJSON(&userAux); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//Comparar se o user e a senha estão corretos

	token, err := service.CreateToken(userAux.IdUser)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
