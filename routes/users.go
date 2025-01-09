package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server.example.com/models"
	"server.example.com/utilities"
)

func signUp(context *gin.Context) {
	var user models.User

	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse the data.",
				"error":   err.Error(),
			})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not save the user.",
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse the data.",
				"error":   err.Error(),
			})
		return
	}

	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized,
			gin.H{
				"message": "Could authenticate the user.",
				"error":   err.Error(),
			})
		return
	}

	token, err := utilities.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not authenticate user.",
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
