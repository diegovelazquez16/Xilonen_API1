package launch

import (

	"log"
	"Xilonen-1/core"

//sensor de humedad:
	sensorHumedadUsecase "Xilonen-1/humedadSuelo/aplication/usecase"
	sensorHumedadRepo "Xilonen-1/humedadSuelo/domain/repository"
	sensorHumedadControllers "Xilonen-1/humedadSuelo/infraestructure/controllers"
	sensorHumedadRoutes "Xilonen-1/humedadSuelo/infraestructure/routes"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"
	"Xilonen-1/websocket"


	"github.com/gin-gonic/gin"
)

func RegisterSensorHumedadModule(router *gin.Engine, sensorHumedadConsumer *sensorHumedadMessaging.SensorHumedadConsumer, wsServer *websocket.WebSocketServer) {
	sensorRepo := &sensorHumedadRepo.SensorHumedadRepositoryImpl{DB: core.GetDB()}

	guardarSensorHumedadUC := &sensorHumedadUsecase.GuardarSensorHumedadUseCase{SensorRepo: sensorRepo}
	obtenerSensoresHumedadUC := &sensorHumedadUsecase.ObtenerSensorHumedadUseCase{SensorRepo: sensorRepo}

	guardarSensorHumedadController := &sensorHumedadControllers.GuardarSensorHumedadController{GuardarSensorHumedadUC: guardarSensorHumedadUC}
	obtenerSensoresHumedadController := &sensorHumedadControllers.ObtenerSensorHumedadController{ObtenerSensorHumedadUC: obtenerSensoresHumedadUC}

	sensorHumedadRoutes.SensorHumedadRoutes(router, guardarSensorHumedadController,obtenerSensoresHumedadController )

	humedadConsumer, err := sensorHumedadMessaging.NewSensorHumedadConsumer(guardarSensorHumedadUC,wsServer)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Humedad Suelo: %v", err)
	}
	go humedadConsumer.Start()
}
