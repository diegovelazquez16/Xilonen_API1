package usecase

import (
	"Xilonen-1/users/domain/models"
	"Xilonen-1/users/domain/repository"
)

type CreateUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *CreateUserUseCase) Execute(user *models.User) error {
	return uc.UserRepo.Create(user)
}