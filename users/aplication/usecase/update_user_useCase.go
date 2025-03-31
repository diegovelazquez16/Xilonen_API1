package usecase

import (
	"Xilonen-1/users/domain/models"
	"Xilonen-1/users/domain/repository"
)

type UpdateUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *UpdateUserUseCase) Execute(user *models.User) error {
	return uc.UserRepo.Update(user)
}