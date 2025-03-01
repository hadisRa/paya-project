package repository

import (
	"paya/models"

	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user *models.User) error
	FindByID(id int) (*models.User, error)
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
	return u.db.Create(user).Error
}

func (u *UserRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}
