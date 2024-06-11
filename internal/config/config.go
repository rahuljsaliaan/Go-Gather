package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	JwtSecret   []byte `envconfig:"JWT_SECRET"`
	JwtExpHours int    `envconfig:"JWT_EXP_HOURS"`
	BcryptCost  int    `envconfig:"BCRYPT_COST"`
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
