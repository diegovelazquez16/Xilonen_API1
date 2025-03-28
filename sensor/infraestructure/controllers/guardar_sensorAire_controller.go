package controllers

import (
	"net/http"
	"Xilonen-1/sensor/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type GuardarSensorController struct {
	GuardarSensorUC *usecase.GuardarSensorUseCase
}

func (sc *GuardarSensorController) GuardarDatos(ctx *gin.Context) {
	var datos struct {
		Valor     float64 `json:"valor"`
		Categoria string `json:"categoria"`

	}

	if err := ctx.ShouldBindJSON(&datos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := sc.GuardarSensorUC.GuardarDatosSensor(datos.Valor, datos.Categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}


