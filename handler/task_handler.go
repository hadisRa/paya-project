package handler

import (
	"net/http"
	"paya/middleware"
	"paya/models"
	"paya/service"
	"strconv"

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
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		task.UserID = ctx.MustGet("user_id").(uint)

		err := h.service.CreateTask(task)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create task",
			})
			return
		}

		middleware.RequestCounterMiddleware(http.StatusOK, ctx.Request.Method, ctx.FullPath())
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"task":   task,
		})
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

func (h *TaskHandler) UpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var task models.Task
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
			return
		}

		err = h.service.UpdateTask(id, &task)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update task"})
			return
		}

		// if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		// 	return
		// }

		//TODO: save in to service
		// database.DB.Save(&task)
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "task": task})
	}
}
func (h *TaskHandler) DeleteTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid"})
			return
		}

		err = h.service.DeleteTask(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete task"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
	}
}
