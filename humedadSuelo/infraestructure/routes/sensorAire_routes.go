package routes

import (
	"Xilonen-1/humedadSuelo/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SensorHumedadRoutes(router *gin.Engine, guardarController *controllers.GuardarSensorHumedadController, obtenerController *controllers.ObtenerSensorHumedadController) {
	group := router.Group("/sensorHumedad")
	{
		group.POST("", guardarController.GuardarDatos)  
		group.GET("", obtenerController.ObtenerDatos)   
	}
}
