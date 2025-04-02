package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"Xilonen-1/users/aplication/usecase"
	"Xilonen-1/users/domain/models"
)

type UserUpdateController struct {
	UpdateUserUC *usecase.UpdateUserUseCase
}

func (c *UserUpdateController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	user.ID = uint(id)

	if err := c.UpdateUserUC.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}