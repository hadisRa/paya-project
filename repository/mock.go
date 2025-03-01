package repository

import (
	"paya/models"

	"github.com/golang/mock/gomock"
)

type MockUserRepository struct {
	*gomock.Mock
}

func (m *MockUserRepository) Create(u *models.User) error {
	return m.Called(u).Errcall()
}

func (m *MockUserRepository) FindByID(id int) (*models.User, error) {
	return m.Called(id).Retval(0).(User), m.Called(id).Errcall()
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	return m.Called(username).Retval(0).(User), m.Called(username).Errcall()
}
