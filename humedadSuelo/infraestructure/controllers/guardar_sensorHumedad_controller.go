package controllers

import (
	"net/http"
	"Xilonen-1/humedadSuelo/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type GuardarSensorHumedadController struct {
	GuardarSensorHumedadUC *usecase.GuardarSensorHumedadUseCase
}

func (sc *GuardarSensorHumedadController) GuardarDatos(ctx *gin.Context) {
	var datos struct {
		ValorHumedad	float64 `json:"valor_humedad"`
		Categoria	    string 	`json:"categoria"`
	}

	if err := ctx.ShouldBindJSON(&datos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := sc.GuardarSensorHumedadUC.GuardarDatosSensorHumedad(datos.ValorHumedad, datos.Categoria); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}


