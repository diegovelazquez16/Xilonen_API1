package launch

import (
	"github.com/gin-gonic/gin"

	"Xilonen-1/sensor/infraestructure/messaging"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"
	sensorNivelAguaMessaging "Xilonen-1/nivelAgua/infraestructure/messaging"
	sensorUVMessaging "Xilonen-1/sensorUV/infraestructure/messaging"
	"Xilonen-1/sensor/infraestructure/websocket"
)


func RegisterRoutes(
	router *gin.Engine,
	sensorAirePublisher *messaging.SensorConsumer,
	sensorHumedadConsumer *sensorHumedadMessaging.SensorHumedadConsumer,
	sensorNivelAguaConsumer *sensorNivelAguaMessaging.SensorNivelAguaConsumer,
	sensorUVConsumer *sensorUVMessaging.SensorUVConsumer,
	wsServer *websocket.WebSocketServer, 
) {

	router.GET("/ws", func(c *gin.Context) {
		wsServer.HandleConnections(c.Writer, c.Request)
	})
	RegisterSensorModule(router, sensorAirePublisher, wsServer)
	RegisterSensorHumedadModule(router, sensorHumedadConsumer)
	RegisterNivelAguaModule(router, sensorNivelAguaConsumer)
	RegisterSensorUVModule(router, sensorUVConsumer)
	RegisterUserModule(router)
}
