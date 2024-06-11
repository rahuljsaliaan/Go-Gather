package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"rahuljsaliaan.com/go-gather/config"
)

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * config.Env.JwtSecretExpire).Unix(),
	})

	return token.SignedString(config.Env.JwtSecret)
}
