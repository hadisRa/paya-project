package main

import (
	"fmt"
	"paya/config"
	"paya/handler"
	"paya/repository"
	"paya/routes"
	"paya/service"

	"github.com/gin-gonic/gin"
)

func main() {
	s := config.LoadConfig()

	taskRepo := repository.NewTaskRepository(s.GormConnection)
	userRepo := repository.NewUserRepository(s.GormConnection)
	//TODO: remove this
	fmt.Println("here", taskRepo, userRepo)

	srv := service.NewTask(taskRepo)
	handr := handler.NewHandlers(srv)

	g := gin.Default()
	routes.Routes(g, handr)

	g.Run(":8080")
}
