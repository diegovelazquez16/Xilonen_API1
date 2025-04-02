package controllers

import (
	"net/http"
	"Xilonen-1/users/aplication/usecase"
	"Xilonen-1/users/domain/models"
	"github.com/gin-gonic/gin"
)

type UserLoginController struct {
	LoginUseCase *usecase.LoginUserUseCase
}

func NewUserLoginController(loginUC *usecase.LoginUserUseCase) *UserLoginController {
	return &UserLoginController{LoginUseCase: loginUC}
}

func (uc *UserLoginController) Login(ctx *gin.Context) {
	var loginData models.LoginRequest

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	token, err := uc.LoginUseCase.Execute(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
