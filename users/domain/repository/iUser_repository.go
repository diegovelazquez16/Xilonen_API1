package repository

import (
	"Xilonen-1/users/domain/models"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	FindByEmail(email string) (*models.User, error)
	GetByEmail(email string) (*models.User, error) 

}
