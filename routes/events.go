package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"server.example.com/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could fetch events. Try again later.",
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse event id.",
				"error":   err.Error(),
			})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch event",
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {

	var event models.Event

	if err := context.ShouldBind(&event); err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse the data.",
				"error":   err.Error(),
			})
		return
	}

	userID := context.GetInt64("userID")
	event.UserID = userID

	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could create event. Try again later.",
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse event id.",
				"error":   err.Error(),
			})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch event",
				"error":   err.Error(),
			})
		return
	}

	userID := context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized,
			gin.H{
				"message": "Not authorized to update event."})
		return
	}

	var UpdatedEvent models.Event

	if err := context.ShouldBind(&UpdatedEvent); err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse the data.",
				"error":   err.Error(),
			})
		return
	}

	UpdatedEvent.ID = eventId

	if err := UpdatedEvent.Update(); err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not update the event.",
				"error":   err.Error(),
			})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Could not parse event id.",
				"error":   err.Error(),
			})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch event",
				"error":   err.Error(),
			})
		return
	}

	userID := context.GetInt64("userID")
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized,
			gin.H{
				"message": "Not authorized to delete event."})
		return
	}

	if err := event.Delete(); err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Could not delete the event.",
				"error":   err,
			})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "The event deleted successfully!"})
}
