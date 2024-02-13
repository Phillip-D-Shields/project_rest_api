package routes

import (
	"net/http"

	"example.com/rest-project/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error parsing request data", "error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error saving user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
