package routes

import (
	"Xilonen-1/sensorTemperatura/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SensorTemperaturaRoutes(router *gin.Engine, guardarController *controllers.GuardarSensorTemperaturaController, obtenerController *controllers.ObtenerSensorTemperaturaController) {
	group := router.Group("/sensorTemperatura")
	{
		group.POST("", guardarController.GuardarDatos)  
		group.GET("", obtenerController.ObtenerDatos)   
	}
}
