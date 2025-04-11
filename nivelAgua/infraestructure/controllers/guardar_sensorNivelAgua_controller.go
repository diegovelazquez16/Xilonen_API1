package controllers

import (
	"net/http"
	"Xilonen-1/nivelAgua/aplication/usecase"
	"github.com/gin-gonic/gin"
)

type GuardarNivelAguaController struct {
	GuardarNivelAguaUC *usecase.GuardarNivelAguaUseCase
}

func (sc *GuardarNivelAguaController) GuardarDatos(ctx *gin.Context) {
	var datos struct {
		NivelAgua     float64 `json:"nivel_agua"`
		Categoria string  `json:"categoria"`  
		Tipo string `json:"tipo"`
	}

	if err := ctx.ShouldBindJSON(&datos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := sc.GuardarNivelAguaUC.GuardarDatosNivelAgua(datos.NivelAgua, datos.Categoria, datos.Tipo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar datos"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}


