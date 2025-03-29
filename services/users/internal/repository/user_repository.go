package repository

import (
	"github.com/titsex/another-messenger/services/users/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{database}
}

func (repository *UserRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User

	err := repository.database.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) CreateUser(user *models.User) error {
	return repository.database.Create(user).Error
}
