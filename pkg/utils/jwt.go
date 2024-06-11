package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"rahuljsaliaan.com/go-gather/internal/config"
)

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * time.Duration(config.Env.JwtExpHours)).Unix(),
	})

	return token.SignedString(config.Env.JwtSecret)
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // assert/check the signing method type

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return config.Env.JwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userId"].(float64))

	return userID, nil
}
