package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/pkg/models"
)

func registerForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")

	if userID == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event. Try again later"})
		return
	}

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	if err = event.Register(userID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event successfully"})
}

func cancelRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")

	if userID == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel Registration. Try again later"})
		return
	}

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	if err = event.CancelRegistration(userID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel Registration. Try again later"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
