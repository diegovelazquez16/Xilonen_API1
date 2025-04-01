package launch

import (

	"log"
	"Xilonen-1/core"

//sensor de Temperatura:
	sensorTemperaturaUsecase "Xilonen-1/sensorTemperatura/aplication/usecase"
	sensorTemperaturaRepo "Xilonen-1/sensorTemperatura/domain/repository"
	sensorTemperaturaControllers "Xilonen-1/sensorTemperatura/infraestructure/controllers"
	sensorTemperaturaRoutes "Xilonen-1/sensorTemperatura/infraestructure/routes"
	sensorTemperaturaMessaging "Xilonen-1/sensorTemperatura/infraestructure/messaging"

	"github.com/gin-gonic/gin"
)

func RegisterSensorTemperaturaModule(router *gin.Engine, sensorTemperaturaConsumer *sensorTemperaturaMessaging.SensorTemperaturaConsumer) {
	sensorRepo := &sensorTemperaturaRepo.SensorTemperaturaRepositoryImpl{DB: core.GetDB()}

	guardarSensorTemperaturaUC := &sensorTemperaturaUsecase.GuardarSensorTemperaturaUseCase{SensorRepo: sensorRepo}
	obtenerSensoresTemperaturaUC := &sensorTemperaturaUsecase.ObtenerSensorTemperaturaUseCase{SensorRepo: sensorRepo}

	guardarSensorTemperaturaController := &sensorTemperaturaControllers.GuardarSensorTemperaturaController{GuardarSensorTemperaturaUC: guardarSensorTemperaturaUC}
	obtenerSensoresTemperaturaController := &sensorTemperaturaControllers.ObtenerSensorTemperaturaController{ObtenerSensorTemperaturaUC: obtenerSensoresTemperaturaUC}

	sensorTemperaturaRoutes.SensorTemperaturaRoutes(router, guardarSensorTemperaturaController,obtenerSensoresTemperaturaController )

	TemperaturaConsumer, err := sensorTemperaturaMessaging.NewSensorTemperaturaConsumer(guardarSensorTemperaturaUC)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Temperatura Suelo: %v", err)
	}
	go TemperaturaConsumer.Start()
}
