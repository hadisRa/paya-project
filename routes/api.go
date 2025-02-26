package routes

import (
	"paya/handler"
	"paya/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine, handler *handler.Handler) {

	g.Group("task")

	g.Use(middleware.AuthMiddleware())
	g.POST("/tasks", handler.TasskHandler.CreateTask())
	 g.GET("/tasks", handler.TasskHandler.GetTask())
	 g.PUT("/tasks/:id", handler.TasskHandler.UpdateTask())
	 g.DELETE("/tasks/:id", handler.TasskHandler.DeleteTask())

	// g.Group("user")
	// g.POST("/register", handler.Register)
	// g.POST("/login", handler.Login)

}
