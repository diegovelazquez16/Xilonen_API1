package launch

import (
	"Xilonen-1/core"
	userUsecase "Xilonen-1/users/aplication/usecase"
	userRepo "Xilonen-1/users/domain/repository"
	userControllers "Xilonen-1/users/infraestructure/controllers"
	userRoutes "Xilonen-1/users/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterUserModule(router *gin.Engine) {
	userRepo := &userRepo.UserRepositoryImpl{DB: core.GetDB()}

	createUserUC := &userUsecase.CreateUserUseCase{UserRepo: userRepo}
	getAllUsersUC := &userUsecase.GetAllUsersUseCase{UserRepo: userRepo}
	getUserUC := &userUsecase.GetUserUseCase{UserRepo: userRepo}
	updateUserUC := &userUsecase.UpdateUserUseCase{UserRepo: userRepo}
	deleteUserUC := &userUsecase.DeleteUserUseCase{UserRepo: userRepo}
	loginUserUC := &userUsecase.LoginUserUseCase{UserRepo: userRepo}

	userCreateController := &userControllers.UserCreateController{CreateUserUC: createUserUC}
	userGetAllController := &userControllers.UserGetAllController{GetAllUsersUC: getAllUsersUC}
	userGetController := &userControllers.UserGetController{GetUserUC: getUserUC}
	userUpdateController := &userControllers.UserUpdateController{UpdateUserUC: updateUserUC}
	userDeleteController := &userControllers.UserDeleteController{DeleteUserUC: deleteUserUC}
	userLoginController := &userControllers.UserLoginController{LoginUseCase:loginUserUC}

	userRoutes.UserRoutes(router, userCreateController, userGetAllController, userUpdateController, userDeleteController, userGetController, userLoginController)
}
