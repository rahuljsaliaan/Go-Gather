package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/pkg/modals"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8000")
}

func getEvents(context *gin.Context) {
	events := modals.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
