package repository

import (
	"paya/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	FindByUserID(userID uint) ([]models.Task, error)
	FindByID(id uint) (*models.Task, error)
	Update(id int,task *models.Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(gdb *gorm.DB) TaskRepository {
	return &taskRepository{
		db: gdb,
	}
}

func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) FindByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Update(id int,task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id int) error {
	return r.db.Delete(id).Error
}
