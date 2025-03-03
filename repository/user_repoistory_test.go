package repository_test

import (
	"testing"

	mock_repository "paya/mock"
	"paya/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserInterface(ctrl)

	user := &models.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
		Password: "password123",
	}

	mockUserRepo.EXPECT().Create(user).Return(nil)
	err := mockUserRepo.Create(user)
	assert.NoError(t, err)
}

func TestUserRepository_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserInterface(ctrl)

	user := &models.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
		Password: "password123",
	}
	mockUserRepo.EXPECT().FindByID(1).Return(user, nil)
	result, err := mockUserRepo.FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_repository.NewMockUserInterface(ctrl)

	user := &models.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
		Password: "password123",
	}
	mockUserRepo.EXPECT().FindByUsername("testuser").Return(user, nil)
	result, err := mockUserRepo.FindByUsername("testuser")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}
