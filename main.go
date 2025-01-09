package main

import (
	"github.com/gin-gonic/gin"
	"server.example.com/db"
	"server.example.com/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
