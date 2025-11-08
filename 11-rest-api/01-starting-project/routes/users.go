package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Error parsing data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Error saving user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"messages": "User created"})
}
