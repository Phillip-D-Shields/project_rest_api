package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/rest-project/db"
	"example.com/rest-project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("kia ora, whanau!")

	db.InitDb()
	fmt.Println("db connected")

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "kia ora, whanau!",
		})
	})

	server.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "thats a healthy api !",
		})
	})

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	log.Fatal(server.Run(":8080"))
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error getting all events", "error": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error creating event", "error": err.Error()})
	}

	// TODO - get user id from token
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error saving event", "error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "event created successfully",
		"event":   event,
	})
}
