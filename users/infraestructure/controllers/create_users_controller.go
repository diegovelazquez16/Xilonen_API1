package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Xilonen-1/users/aplication/usecase"
	"Xilonen-1/users/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type UserCreateController struct {
	CreateUserUC *usecase.CreateUserUseCase
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (c *UserCreateController) Create(ctx *gin.Context) {
	var user models.User
	
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email y contraseña son requeridos"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
		return
	}
	user.Password = string(hashedPassword)

	err = c.CreateUserUC.Execute(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Nombre,
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user":    response,
	})
}