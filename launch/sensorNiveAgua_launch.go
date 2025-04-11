// launch/sensorNivelAgua_launch.go
package launch

import (
	"log"
	"Xilonen-1/core"

//sensor de nivel de agua:
	nivelAguaUsecase "Xilonen-1/nivelAgua/aplication/usecase"
	nivelAguaRepo "Xilonen-1/nivelAgua/domain/repository"
	nivelAguaControllers "Xilonen-1/nivelAgua/infraestructure/controllers"
	nivelAguaRoutes "Xilonen-1/nivelAgua/infraestructure/routes"
	sensorNivelAguaMessaging "Xilonen-1/nivelAgua/infraestructure/messaging"
	"Xilonen-1/websocket"



	"github.com/gin-gonic/gin"
)

func RegisterNivelAguaModule(router *gin.Engine, sensorNivelAguaConsumer * sensorNivelAguaMessaging.SensorNivelAguaConsumer, wsServer *websocket.WebSocketServer) {
	nivelAguaRepo := &nivelAguaRepo.NivelAguaRepositoryImpl{DB: core.GetDB()}

	guardarNivelAguaUC := &nivelAguaUsecase.GuardarNivelAguaUseCase{NivelAguaRepo: nivelAguaRepo}
	obtenerNivelAguaUC := &nivelAguaUsecase.ObtenerNivelAguaUseCase{NivelAguaRepo: nivelAguaRepo}

	guardarNivelAguaController := &nivelAguaControllers.GuardarNivelAguaController{GuardarNivelAguaUC: guardarNivelAguaUC}
	obtenerNivelAguaController := &nivelAguaControllers.ObtenerNivelAguaController{ObtenerNivelAguaUC: obtenerNivelAguaUC}

	nivelAguaRoutes.NivelAguaRoutes(router, guardarNivelAguaController,obtenerNivelAguaController )
	
	nivelAguaConsumer, err := sensorNivelAguaMessaging.NewSensorNivelAguaConsumer(guardarNivelAguaUC, wsServer)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor T1592: %v", err)
	}

	go nivelAguaConsumer.Start()
}
