package controllers

import (
	"net/http"
	"holamundo/sensor/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type GuardarSensorController struct {
	GuardarSensorUC *usecase.GuardarSensorUseCase
}

func (sc *GuardarSensorController) GuardarDatos(ctx *gin.Context) {
	var datos struct {
		CO2     float64 `json:"co2"`
		NH3     float64 `json:"nh3"`
		Alcohol float64 `json:"alcohol"`
		Tolueno float64 `json:"tolueno"`
		Acetona float64 `json:"acetona"`
	}

	if err := ctx.ShouldBindJSON(&datos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := sc.GuardarSensorUC.GuardarDatosSensor(datos.CO2, datos.NH3, datos.Alcohol, datos.Tolueno, datos.Acetona); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}


