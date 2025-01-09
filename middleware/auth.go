package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server.example.com/utilities"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Not Unauthorized"})
		return
	}

	userID, err := utilities.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Not Unauthorized"})
		return
	}

	context.Set("userID", userID)
	context.Next()

}
