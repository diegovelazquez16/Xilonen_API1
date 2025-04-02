package routes

import (
	"Xilonen-1/nivelAgua/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func NivelAguaRoutes(router *gin.Engine, guardarController *controllers.GuardarNivelAguaController, obtenerController *controllers.ObtenerNivelAguaController) {
	group := router.Group("/nivelAgua")
	{
		group.POST("", guardarController.GuardarDatos)  
		group.GET("", obtenerController.ObtenerDatos)   
	}
}
