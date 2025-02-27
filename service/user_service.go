package service

import (
	"paya/models"
	"paya/repository"
)

type User interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
}

type UserService struct {
	Repo repository.UserRepository
}

func NewUser(repo repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (t *UserService) CreateUser(user *models.User) error {
	err := t.Repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
func (t *UserService) GetUser(id int) (*models.User, error) {
	user, err := t.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
