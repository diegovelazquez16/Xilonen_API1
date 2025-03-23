package controllers

import (
	"net/http"
	"Xilonen-1/humedadSuelo/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type ObtenerSensorHumedadController struct {
	ObtenerSensorHumedadUC *usecase.ObtenerSensorHumedadUseCase
}

func (sc *ObtenerSensorHumedadController) ObtenerDatos(ctx *gin.Context) {
	sensoresHumedad, err := sc.ObtenerSensorHumedadUC.ObtenerDatosSensores()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}

	ctx.JSON(http.StatusOK, sensoresHumedad)
}