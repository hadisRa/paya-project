package handler

import (
	"errors"
	"net/http"
	"paya/middleware"
	"paya/models"
	"paya/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
				"status":  http.StatusBadRequest,
				"error":   "@handler.task_handler.CreateTask",
				"message": err.Error(),
			})
			return
		}

		userId, ok := ctx.Get("user_id")
		if !ok {
			middleware.RequestCounterMiddleware(http.StatusUnauthorized, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "@handler.task_handler.CreateTask",
				"message": "User ID not found in the context",
			})
			return
		}

		userIdUint, ok := userId.(uint)
		if !ok {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.task_handler.CreateTask",
				"message": "Invalid user ID type",
			})
			return
		}

		task.UserID = userIdUint

		err := h.service.CreateTaskByUserID(task)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.task_handler.CreateTask.CreateTaskByUserID",
				"message": "Failed to create task",
			})
			return
		}

		middleware.RequestCounterMiddleware(http.StatusOK, ctx.Request.Method, ctx.FullPath())
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	}
}

func (h *TaskHandler) GetUserTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, ok := ctx.Get("user_id")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "@handler.task_handler.GetUserTask",
				"message": "User ID not found in the context",
			})
			return
		}

		userIdUint, ok := userId.(uint)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.task_handler.GetUserTask",
				"message": "Invalid user ID type",
			})
			return
		}

		tasks, err := h.service.GetUserTaskByUserID(userIdUint)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.task_handler.GetUserTask.GetUserTaskByUserID",
				"message": "Failed to Get user tasks",
			})
			return
		}

		middleware.RequestCounterMiddleware(http.StatusOK, ctx.Request.Method, ctx.FullPath())
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"task":   tasks,
		})
	}
}

func (h *TaskHandler) UpdateTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, ok := ctx.Get("id")
		if !ok {
			middleware.RequestCounterMiddleware(http.StatusUnauthorized, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "@handler.task_handler.UpdateTask",
				"message": "ID not found in the context",
			})
			return
		}

		ID, ok := id.(int)
		if !ok {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.task_handler.UpdateTask",
				"message": "Invalid ID type",
			})
			return
		}

		var task models.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.task_handler.UpdateTask",
				"message": "Invalid request body",
			})
			return
		}

		err := h.service.UpdateTaskByID(ID, &task)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				middleware.RequestCounterMiddleware(http.StatusNotFound, ctx.Request.Method, ctx.FullPath())
				ctx.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"error":   "@handler.task_handler.UpdateTask.UpdateTaskByID",
					"message": "Task not found",
				})
				return
			}
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.task_handler.UpdateTask.UpdateTaskByID",
				"message": "Failed to update task",
			})
			return
		}

		middleware.RequestCounterMiddleware(http.StatusOK, ctx.Request.Method, ctx.FullPath())
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"task":   task,
		})
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
