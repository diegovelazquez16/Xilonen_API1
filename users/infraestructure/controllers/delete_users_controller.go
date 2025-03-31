package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"Xilonen-1/users/aplication/usecase"
)

type UserDeleteController struct {
	DeleteUserUC *usecase.DeleteUserUseCase
}

func (c *UserDeleteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.DeleteUserUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}