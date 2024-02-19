package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error fetching event id", "error": err.Error(), "id": eventId})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting event by id", "error": err.Error(), "id": eventId})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error registering for event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "registered for event successfully",
		"event":   event,
	})
}

func unregisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error fetching event id", "error": err.Error(), "id": eventId})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting event by id", "error": err.Error(), "id": eventId})
		return
	}

	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error unregistering for event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "unregistered for event successfully",
		"event":   event,
	})
}
