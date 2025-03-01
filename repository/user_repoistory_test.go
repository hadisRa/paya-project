package repository

import (
	"paya/models"
	. "paya/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Mock structs
type MockDB struct {
	gomock.Controller
}

func (m *MockDB) Create(user *models.User) error {
	return nil
}

func (m *MockDB) Where(query string, args ...interface{}) *gorm.DB {
	return m.DB()
}

func (m *MockDB) First(dest interface{}, query ...interface{}) *gorm.DB {
	return m.DB()
}

func TestUserRepository(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockDB(ctrl)
	repo := NewUserRepository(mockDB)

	// Test Create method
	user := &User{Name: "Test User"}
	mockDB.EXPECT().Create(gomock.Any()).Return(nil)

	err := repo.Create(user)
	assert.NoError(t, err)

	// Test FindByID method
	mockDB.EXPECT().Where("id = ?", 1).First(gomock.Any()).Return(nil)

	foundUser, _ := repo.FindByID(1)
	assert.NotNil(t, foundUser)

	// Test FindByUsername method
	mockDB.EXPECT().Where("username = ?", "testuser").First(gomock.Any()).Return(nil)

	foundUser, _ = repo.FindByUsername("testuser")
	assert.NotNil(t, foundUser)
}
