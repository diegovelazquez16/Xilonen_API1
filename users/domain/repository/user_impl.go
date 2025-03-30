package repository

import (
	"gorm.io/gorm"
	"errors"
	"Xilonen-1/users/domain/models"
)


type UserRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}
func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) Update(user *models.User) error{
	return r.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
} 

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("usuario no encontrado")
	}
	return &user, nil
}



func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
