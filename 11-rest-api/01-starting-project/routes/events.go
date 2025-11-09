package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching events"})
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

	e.UserID = ctx.GetInt64("userId")
	err = e.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Errorsaving event"})
		return
	}

	ctx.JSON(http.StatusCreated, e)
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID must be an integer"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID must be an integer"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	userId := ctx.GetInt64("userId")
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to update this event"})
		return
	}

	var e models.Event
	err = ctx.ShouldBindJSON(&e)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing request"})
		return
	}

	e.ID = id
	err = e.Update()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Error updating event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID must be an integer"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
		return
	}

	userId := ctx.GetInt64("userId")
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Error deleting event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
