package handler

import (
	"net/http"
	"paya/models"
	"paya/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service service.Task
}

func NewTaskHandler(srv service.Task) *TaskHandler {
	return &TaskHandler{
		service: srv,
	}
}

func (h *TaskHandler) CreateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var task models.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task.UserID = ctx.MustGet("user_id").(uint)

		err := h.service.CreateTask(task)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "task": task})
	}
}

func (h *TaskHandler) GetTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId := ctx.MustGet("user_id").(uint)
		//TODO: check the id or user id?
		task, err := h.service.GetTask(userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get task"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "task": task})

	}
}

func UpdateTask(c *gin.Context) {
	// var task models.Task
	// id := c.Param("id")

	// if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	// 	return
	// }

	// if err := c.ShouldBindJSON(&task); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// database.DB.Save(&task)
	// c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	// id := c.Param("id")
	// var task models.Task

	// if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	// 	return
	// }

	// database.DB.Delete(&task)
	// c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
