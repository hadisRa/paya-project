package repository

import (
	"paya/models"

	"gorm.io/gorm"
)

type TaskInterface interface {
	Create(task *models.Task) error
	Find(id int, task *models.Task) error
	FindByUserID(userID uint) ([]models.Task, error)
	FindByID(id uint) (*models.Task, error)
	Update(id int, task *models.Task) error
	Delete(id int) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(gdb *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: gdb,
	}
}

func (t *TaskRepository) Create(task *models.Task) error {
	return t.db.Create(task).Error
}

func (t *TaskRepository) Find(id int, task *models.Task) error {
	return t.db.First(task, id).Error
}

func (t *TaskRepository) FindByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := t.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (t *TaskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	err := t.db.First(&task, id).Error
	return &task, err
}

func (t *TaskRepository) Update(id int, task *models.Task) error {
	return t.db.Model(&models.Task{}).Where("id = ?", id).Updates(task).Error
}

func (t *TaskRepository) Delete(id int) error {
	return t.db.Delete(id).Error
}
