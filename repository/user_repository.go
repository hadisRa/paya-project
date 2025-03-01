package repository

import (
	"paya/models"

	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user *models.User) error
	FindByID(id int) (*models.User, error)
	FindByUsername(userName string) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *models.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FindByUsername(userName string) (*models.User, error) {
	var user models.User
	err := u.db.Where("username = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
