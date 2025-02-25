package routes

import (
	"paya/controllers"
	"paya/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetTasks)
	router.PUT("/tasks/:id", controllers.UpdateTask)   // مسیر به‌روزرسانی تسک
    router.DELETE("/tasks/:id", controllers.DeleteTask) // مسیر حذف تسک
}
