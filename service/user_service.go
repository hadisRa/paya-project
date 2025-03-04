package service

import (
	"paya/models"
	"paya/repository"
)

type User interface {
	CreateUser(user *models.User) error
	GetUser(userName string) (*models.User, error)
}

type UserService struct {
	Repo repository.UserInterface
}

func NewUser(repo repository.UserInterface) *UserService {
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
func (t *UserService) GetUser(userName string) (*models.User, error) {
	user, err := t.Repo.FindByUsername(userName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
