package routes

import (
	"Xilonen-1/sensorUV/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func SensorUVRoutes(router *gin.Engine, guardarController *controllers.GuardarSensorUVController, obtenerController *controllers.ObtenerSensorUVController) {
	group := router.Group("/sensorUV")
	{
		group.POST("", guardarController.GuardarDatos)  
		group.GET("", obtenerController.ObtenerDatos)   
	}
}
