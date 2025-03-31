package usecase

import (
	"Xilonen-1/users/domain/models"
	"Xilonen-1/users/domain/repository"
)
type GetUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *GetUserUseCase) Execute(id uint) (*models.User, error) {
	return uc.UserRepo.GetByID(id)
}