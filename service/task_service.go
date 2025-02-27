package service

import (
	"paya/models"
	"paya/repository"
)

type Task interface {
	CreateTaskByUserID(item models.Task) error
	GetUserTaskByUserID(userID uint) ([]models.Task, error)
	UpdateTaskByID(id int, item *models.Task) error
	DeleteTask(id int) error
}

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTask(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (t *TaskService) CreateTaskByUserID(item models.Task) error {
	err := t.Repo.Create(&item)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskService) GetUserTaskByUserID(userID uint) ([]models.Task, error) {
	task, err := t.Repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskService) UpdateTaskByID(id int, task *models.Task) error {
	existingTask := &models.Task{}
	if err := t.Repo.Find(id, existingTask); err != nil {
		return err
	}

	return t.Repo.Update(id, task)
}

func (t *TaskService) DeleteTask(id int) error {
	err := t.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
