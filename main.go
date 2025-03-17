package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prezessikora/events/db"
	"github.com/prezessikora/events/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
