package launch

import (
	"log"
	"Xilonen-1/core"
	sensorUsecase "Xilonen-1/sensor/aplication/usecase"
	sensorRepo "Xilonen-1/sensor/domain/repository"
	sensorControllers "Xilonen-1/sensor/infraestructure/controllers"
	sensorRoutes "Xilonen-1/sensor/infraestructure/routes"
	sensorMessaging "Xilonen-1/sensor/infraestructure/messaging"

	"github.com/gin-gonic/gin"
)

func RegisterSensorModule(router *gin.Engine, sensorPublisher *sensorMessaging.SensorConsumer) {
	sensorRepo := &sensorRepo.SensorRepositoryImpl{DB: core.GetDB()}

	guardarSensorUC := &sensorUsecase.GuardarSensorUseCase{SensorRepo: sensorRepo}
	obtenerSensoresUC := &sensorUsecase.ObtenerSensorUseCase{SensorRepo: sensorRepo}

	guardarSensorController := &sensorControllers.GuardarSensorController{GuardarSensorUC: guardarSensorUC}
	obtenerSensoreController := &sensorControllers.ObtenerSensorController{ObtenerSensorUC: obtenerSensoresUC}

	sensorRoutes.SensorRoutes(router, guardarSensorController,obtenerSensoreController )

	sensorConsumer, err := sensorMessaging.NewSensorConsumer(&sensorUsecase.GuardarSensorUseCase{})
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para DHT11: %v", err)
	}
	go sensorConsumer.Start()
}
