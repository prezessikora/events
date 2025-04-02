package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prezessikora/events/db"
	"github.com/prezessikora/events/routes"
)

func main() {

	db.InitDB()
	v := 10
	
	server := gin.Default()
	routes.RegisterRoutes(server)
	fmt.Sprintf("aaaa")
	server.Run(":8080")
	j := 0
	for i := 0; i < 10; j++ {

	}
}
