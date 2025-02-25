package controllers

import (
	"net/http"
	"paya/database"
	"paya/models"

	"github.com/gin-gonic/gin"
)

// CreateTask ایجاد یک تسک جدید
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.UserID=c.MustGet("user_id").(uint)
	database.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

// GetTasks دریافت تمامی تسک‌ها
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	userId:=c.MustGet("user_id").(uint)
	database.DB.Where("user_id =?",userId).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// UpdateTask به‌روزرسانی یک تسک بر اساس شناسه آن
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	// پیدا کردن تسک بر اساس شناسه
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// بایند کردن داده‌های جدید
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// به‌روزرسانی تسک
	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask حذف یک تسک بر اساس شناسه آن
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// پیدا کردن تسک بر اساس شناسه
	if err := database.DB.Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// حذف تسک
	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
