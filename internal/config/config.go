package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	JwtSecret       []byte        `envconfig:"JWT_SECRET"`
	JwtSecretExpire time.Duration `envconfig:"JWT_SECRET_EXPIRE"`
	BcryptCost      int           `envconfig:"BCRYPT_COST"`
}

var Env env

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := envconfig.Process("", &Env); err != nil {
		log.Fatal(err.Error())
	}
}
