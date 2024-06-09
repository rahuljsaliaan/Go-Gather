package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/internal/db"
	"rahuljsaliaan.com/go-gather/pkg/models"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvent)

	server.Run(":8000")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEventByID(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format"})
		return
	}

	event.ID = 1
	event.UserId = 1

	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
