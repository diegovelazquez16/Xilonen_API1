package controllers

import (
	"net/http"
	"Xilonen-1/sensorUV/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type GuardarSensorUVController struct {
	GuardarSensorUVUC *usecase.GuardarSensorUVUseCase
}

func (sc *GuardarSensorUVController) GuardarDatos(ctx *gin.Context) {
	var datos struct {
		ValorUV     float64 `json:"valor_uv"`//OJO
		Categoria string  `json:"categoria"`  


	}

	if err := ctx.ShouldBindJSON(&datos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := sc.GuardarSensorUVUC.GuardarDatosSensorUV(datos.ValorUV, datos.Categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}


