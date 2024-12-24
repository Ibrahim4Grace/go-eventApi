package main

import (
	"event/db"
	"event/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// start my server
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	server := gin.Default()

	err = server.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	// RegisterRoutes is a function that takes a server and registers all the routes
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
