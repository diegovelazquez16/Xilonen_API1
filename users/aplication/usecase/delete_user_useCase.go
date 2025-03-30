package usecase

import (
	"Xilonen-1/users/domain/repository"
)
type DeleteUserUseCase struct {
	UserRepo repository.IUserRepository
}

func (uc *DeleteUserUseCase) Execute(id uint) error {
	return uc.UserRepo.Delete(id)
}