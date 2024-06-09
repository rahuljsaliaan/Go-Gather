package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/internal/db"
	"rahuljsaliaan.com/go-gather/pkg/models"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.POST("/events", createEvent)
	server.GET("/events", getEvents)

	server.Run(":8000")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		log.Printf("Error binding JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format"})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
