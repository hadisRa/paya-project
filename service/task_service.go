package service

import (
	"paya/models"
	"paya/repository"
)

type Task interface {
	CreateTask(item models.Task) error
	GetTask(id uint) (*models.Task, error)
}

type TaskService struct {
	Repo repository.TaskRepository
}

func NewTask(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (t *TaskService) CreateTask(item models.Task) error {
	err := t.Repo.Create(&item)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskService) GetTask(id uint) (*models.Task, error) {
	task, err := t.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}
