package routes

import (
	"event/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	//YOU CAN USE
	// server.POST("/events", middlewares.Authenticate, createEvents)
	// server.PUT("/events/:id",middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate,deleteEvent)

	//OR
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
