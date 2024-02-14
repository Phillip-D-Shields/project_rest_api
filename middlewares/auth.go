package middlewares

import (
	"net/http"

	"example.com/rest-project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	tokem := context.Request.Header.Get("Authorization")

	if tokem == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userId, err := utils.VerifyToken(tokem)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "error verifying token", "error": err.Error()})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
