package controllers

import (
	"net/http"
	"Xilonen-1/sensor/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type ObtenerSensorController struct {
	ObtenerSensorUC *usecase.ObtenerSensorUseCase
}

func (sc *ObtenerSensorController) ObtenerDatos(ctx *gin.Context) {
	sensores, err := sc.ObtenerSensorUC.ObtenerDatosSensores()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}

	ctx.JSON(http.StatusOK, sensores)
}