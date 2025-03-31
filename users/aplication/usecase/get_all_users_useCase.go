package usecase

import (
	"Xilonen-1/users/domain/models"
	"Xilonen-1/users/domain/repository"
)
type GetAllUsersUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *GetAllUsersUseCase) Execute() ([]models.User, error) {
	return uc.UserRepo.GetAll()
}

