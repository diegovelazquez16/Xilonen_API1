package controllers

import (
	"net/http"
	"Xilonen-1/nivelAgua/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type ObtenerNivelAguaController struct {
	ObtenerNivelAguaUC *usecase.ObtenerNivelAguaUseCase
}

func (sc *ObtenerNivelAguaController) ObtenerDatos(ctx *gin.Context) {
	sensores, err := sc.ObtenerNivelAguaUC.ObtenerNivelAgua()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}

	ctx.JSON(http.StatusOK, sensores)
}