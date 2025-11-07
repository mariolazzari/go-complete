package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", postEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching evnts"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func postEvent(ctx *gin.Context) {
	var e models.Event

	err := ctx.ShouldBindJSON(&e)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}

	e.ID = 1
	e.UserID = 1
	err = e.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Errorsaving event"})
		return
	}

	ctx.JSON(http.StatusCreated, e)
}
