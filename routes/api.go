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
	t.POST("/", middleware.RateLimiter(10, 20), handler.TasskHandler.CreateTask())
	t.GET("/", handler.TasskHandler.GetUserTask())
	t.PUT("/Update/:id", handler.TasskHandler.UpdateTask())
	t.DELETE("/Delete/:id", handler.TasskHandler.DeleteTask())

	u := g.Group("user")
	u.POST("/register", handler.UserHandler.Register())
	u.POST("/login", handler.UserHandler.Login())

}
