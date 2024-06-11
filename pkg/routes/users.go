package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/pkg/models"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format"})
		return
	}

	err := user.Save()

	if err != nil {
		log.Printf("%v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not sign up. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Signed up successfully"})
}
