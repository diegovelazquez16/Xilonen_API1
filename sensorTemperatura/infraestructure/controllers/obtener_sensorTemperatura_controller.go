package controllers

import (
	"net/http"
	"Xilonen-1/sensorTemperatura/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type ObtenerSensorTemperaturaController struct {
	ObtenerSensorTemperaturaUC *usecase.ObtenerSensorTemperaturaUseCase
}

func (sc *ObtenerSensorTemperaturaController) ObtenerDatos(ctx *gin.Context) {
	sensoresTemperatura, err := sc.ObtenerSensorTemperaturaUC.ObtenerDatosSensoresTemperatura()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}

	ctx.JSON(http.StatusOK, sensoresTemperatura)
}