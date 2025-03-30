package launch

import (
	"log"
	"Xilonen-1/core"

//sensor de UV:
	sensorUVUsecase "Xilonen-1/sensorUV/aplication/usecase"
	sensorUVRepo "Xilonen-1/sensorUV/domain/repository"
	sensorUVControllers "Xilonen-1/sensorUV/infraestructure/controllers"
	sensorUVRoutes "Xilonen-1/sensorUV/infraestructure/routes"
	sensorUVMessaging "Xilonen-1/sensorUV/infraestructure/messaging"

	"github.com/gin-gonic/gin"
)

func RegisterSensorUVModule(router *gin.Engine, sensorUVPublisher * sensorUVMessaging.SensorUVConsumer) {
	sensorUVRepo := &sensorUVRepo.SensorUVRepositoryImpl{DB: core.GetDB()}

	guardarSensorUVUC := &sensorUVUsecase.GuardarSensorUVUseCase{SensorUVRepo: sensorUVRepo}
	obtenerSensoresUVUC := &sensorUVUsecase.ObtenerSensorUVUseCase{SensorUVRepo: sensorUVRepo}

	guardarSensorUVController := &sensorUVControllers.GuardarSensorUVController{GuardarSensorUVUC: guardarSensorUVUC}
	obtenerSensoresUVController := &sensorUVControllers.ObtenerSensorUVController{ObtenerSensorUVUC: obtenerSensoresUVUC}

	sensorUVRoutes.SensorUVRoutes(router, guardarSensorUVController,obtenerSensoresUVController )

	uvConsumer, err := sensorUVMessaging.NewSensorUVConsumer(guardarSensorUVUC)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Humedad Suelo: %v", err)
	}
	go uvConsumer.Start()
}
//UV