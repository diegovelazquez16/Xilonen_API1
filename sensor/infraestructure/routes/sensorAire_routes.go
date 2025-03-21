package routes

import (
	"holamundo/sensor/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SensorRoutes(router *gin.Engine, guardarController *controllers.GuardarSensorController, obtenerController *controllers.ObtenerSensorController) {
	group := router.Group("/sensorAire")
	{
		group.POST("", guardarController.GuardarDatos)  
		group.GET("", obtenerController.ObtenerDatos)   
	}
}
