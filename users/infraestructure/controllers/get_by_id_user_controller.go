package controllers


import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"Xilonen-1/users/aplication/usecase"
)

type UserGetController struct {
	GetUserUC *usecase.GetUserUseCase
}

func (c *UserGetController) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe estar registrado o ser un numero valido"})
		return
	}

	user, err := c.GetUserUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado."})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
