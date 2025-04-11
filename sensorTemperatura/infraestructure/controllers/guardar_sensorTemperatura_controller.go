package controllers

import (
	"net/http"
	"Xilonen-1/sensorTemperatura/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type GuardarSensorTemperaturaController struct {
	GuardarSensorTemperaturaUC *usecase.GuardarSensorTemperaturaUseCase
}

func (sc *GuardarSensorTemperaturaController) GuardarDatos(ctx *gin.Context) {
	var datos struct {
		ID 				uint 	`json:"id"`
		ValorTemperatura	float64 `json:"valor_Temperatura"`
		Categoria	    string 	`json:"categoria"`
		Tipo string `json:"tipo"`

	}

	if err := ctx.ShouldBindJSON(&datos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := sc.GuardarSensorTemperaturaUC.GuardarDatosSensorTemperatura(datos.ID, datos.ValorTemperatura, datos.Categoria, datos.Tipo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}


