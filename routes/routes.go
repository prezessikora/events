package routes

import (
	"com.sikora/events/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById) //events/1 events/2
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	server.POST("/users", createUser)
	server.POST("/login", login)
}
