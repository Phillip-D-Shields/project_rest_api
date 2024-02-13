package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// ! health check routes =========================
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
	// ! event routes =========================
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// ! user routes =========================
	server.POST("/signup", createUser)
	server.POST("/login", loginUser)
}
