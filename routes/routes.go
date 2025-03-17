package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prezessikora/events/middleware"
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
