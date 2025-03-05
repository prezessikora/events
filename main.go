package main

import (
	"com.sikora/events/db"
	"com.sikora/events/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
