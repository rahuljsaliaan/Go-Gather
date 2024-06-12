package routes

import (
	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/pkg/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// authenticated
	authenticated := server.Group("/", middlewares.Authenticate)

	// user
	server.POST("/signup", signup)
	server.POST("/login", login)

	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// registration
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
