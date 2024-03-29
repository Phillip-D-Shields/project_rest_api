package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-project/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting all events", "error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error fetching id", "error": err.Error(), "id": id})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting event by id", "error": err.Error(), "id": id})
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error creating event", "error": err.Error()})
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error saving event", "error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "event created successfully",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error fetching event id", "error": err.Error(), "id": id})
		return
	}

	userID := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting event by id", "error": err.Error(), "id": id})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error binding jSON for event", "error": err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error updating event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "event updated successfully",
		"event":   updatedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error fetching event id", "error": err.Error(), "id": id})
		return
	}

	userID := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting event by id", "error": err.Error(), "id": id})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not allowed to update this event"})
		return
	}

	err = event.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error deleting event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "event deleted successfully",
		"id":      id,
	})
}
