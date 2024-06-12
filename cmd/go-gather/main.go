package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"rahuljsaliaan.com/go-gather/internal/config"
	"rahuljsaliaan.com/go-gather/internal/db"
	"rahuljsaliaan.com/go-gather/pkg/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db.InitDB()
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	routes.RegisterRoutes(server)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Env.Port), server))
}
