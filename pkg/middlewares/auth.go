package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/pkg/utils"
)

func Authenticate(context *gin.Context) {
	cookie, err := context.Request.Cookie("token")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not authorized"})
		return
	}

	token := cookie.Value

	userID, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userID", userID)

	context.Next()
}
