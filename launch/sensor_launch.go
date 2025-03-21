package launch

import (
	"holamundo/core"
	sensorUsecase "holamundo/sensor/aplication/usecase"
	sensorRepo "holamundo/sensor/domain/repository"
	sensorControllers "holamundo/sensor/infraestructure/controllers"
	sensorRoutes "holamundo/sensor/infraestructure/routes"
	"holamundo/sensor/infraestructure/messaging"

	"github.com/gin-gonic/gin"
)

func RegisterSensorModule(router *gin.Engine, sensorPublisher *messaging.SensorPublisher) {
	sensorRepo := &sensorRepo.SensorRepositoryImpl{DB: core.GetDB()}

	guardarSensorUC := &sensorUsecase.GuardarSensorUseCase{SensorRepo: sensorRepo}
	obtenerSensoresUC := &sensorUsecase.ObtenerSensorUseCase{SensorRepo: sensorRepo}

	guardarSensorController := &sensorControllers.GuardarSensorController{GuardarSensorUC: guardarSensorUC}
	obtenerSensoreController := &sensorControllers.ObtenerSensorController{ObtenerSensorUC: obtenerSensoresUC}

	sensorRoutes.SensorRoutes(router, guardarSensorController,obtenerSensoreController )
}
