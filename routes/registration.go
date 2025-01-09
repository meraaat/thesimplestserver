package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"server.example.com/models"
)

func registerForEvent(context *gin.Context) {

	userID := context.GetInt64("userID")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse the event id.",
				"error":   err,
			})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch the event.",
				"error":   err,
			})
		return
	}

	if err := event.Registration(userID); err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not register user for event.",
				"error":   err,
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered for the event successfully"})
}

func cancelRegistration(context *gin.Context) {

	userID := context.GetInt64("userID")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse the event id.",
				"error":   err,
			})
		return
	}

	var event models.Event
	event.ID = eventId

	if err := event.CancelRegistration(userID); err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not cancel registration.",
				"error":   err,
			})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled successfully"})

}
