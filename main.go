package main

import (
	"event/db"
	"event/models"
	"net/http"
	"strconv"

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

	// GET, POST, PUT, PATCH, DDLEETE
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	//getting all events
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	// .JSON return json content back
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {
	//getting event by id
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GeteventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {
	//create  events from the model struct
	var event models.Event

	// .JSON return json content back
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cound not parse request data"})
		return
	}

	//user dummy value to save the id for now
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later."})
		return
	}

	//send back value if everything work correctly
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
