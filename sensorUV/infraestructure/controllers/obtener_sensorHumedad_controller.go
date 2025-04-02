package controllers

import (
	"net/http"
	"Xilonen-1/sensorUV/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type ObtenerSensorUVController struct {
	ObtenerSensorUVUC *usecase.ObtenerSensorUVUseCase
}

func (sc *ObtenerSensorUVController) ObtenerDatos(ctx *gin.Context) {
	sensoresUV, err := sc.ObtenerSensorUVUC.ObtenerDatosSensoresUV()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}

	ctx.JSON(http.StatusOK, sensoresUV)
}