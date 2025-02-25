package models

import "github.com/jinzhu/gorm"

type Task struct {
    gorm.Model
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"` // To-Do, In Progress, Done
    UserID      uint   `json:"user_id"`
}
