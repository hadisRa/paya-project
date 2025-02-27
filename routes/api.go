package routes

import (
	"paya/handler"
	"paya/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine, handler *handler.Handler) {

	g.Use(middleware.AuthMiddleware())
	g.Use(middleware.RequestDurationMiddleware())
	g.Use(middleware.RateLimiter(5, 10))

	t := g.Group("task")
	t.POST("/", middleware.RateLimiter(10, 20), handler.TaskHandler.CreateTask())
	t.GET("/", handler.TaskHandler.GetUserTask())
	t.PUT("/update/:id", handler.TaskHandler.UpdateTask())
	t.DELETE("/delete/:id", handler.TaskHandler.DeleteTask())

	u := g.Group("user")
	u.POST("/register", handler.UserHandler.Register())
	u.POST("/login", handler.UserHandler.Login())

}
