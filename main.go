package main

import (
	"event/db"
	"event/routes"

	"github.com/gin-gonic/gin"
)

// start my server
func main() {
	db.InitDB()
	server := gin.Default()

	err := server.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	// RegisterRoutes is a function that takes a server and registers all the routes
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
