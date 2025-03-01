package main

import (
	"paya/config"
	"paya/handler"
	"paya/metrics"
	"paya/repository"
	"paya/routes"
	"paya/service"

	"github.com/gin-gonic/gin"
)

func main() {
	s := config.LoadConfig()

	taskRepo := repository.NewTaskRepository(s.GormConnection)
	userRepo := repository.NewUserRepository(s.GormConnection)
	cacheRepo := repository.NewCacheRepository(s.RedisConnection)

	taskSrv := service.NewTask(taskRepo)
	userSrv := service.NewUser(userRepo)

	handr := handler.NewHandlers(taskSrv, userSrv, cacheRepo, userRepo)

	g := gin.Default()
	g.GET("/metrics", gin.WrapH(metrics.MetricsHandler()))

	routes.Routes(g, handr)

	g.Run(":8080")
}
