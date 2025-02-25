package main

import (
	"paya/config"
	"paya/database"
	"paya/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.Connect()
	router := gin.Default()
	routes.AuthRoutes(router)
	routes.TaskRoutes(router)
	router.Run(":8080")
}
