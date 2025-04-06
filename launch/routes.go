package launch

import (
	"github.com/gin-gonic/gin"

	"Xilonen-1/sensor/infraestructure/messaging"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"
	sensorNivelAguaMessaging "Xilonen-1/nivelAgua/infraestructure/messaging"
	sensorUVMessaging "Xilonen-1/sensorUV/infraestructure/messaging"
	sensorTemperaturaMessaging "Xilonen-1/sensorTemperatura/infraestructure/messaging"
	"Xilonen-1/sensor/infraestructure/websocket"
)


func RegisterRoutes(
	router *gin.Engine,
	sensorAirePublisher *messaging.SensorConsumer,
	sensorHumedadConsumer *sensorHumedadMessaging.SensorHumedadConsumer,
	sensorNivelAguaConsumer *sensorNivelAguaMessaging.SensorNivelAguaConsumer,
	sensorUVConsumer *sensorUVMessaging.SensorUVConsumer,
	SensorTemperaturaConsumer * sensorTemperaturaMessaging.SensorTemperaturaConsumer,
	wsServer *websocket.WebSocketServer, 
) {

	router.GET("/ws", func(c *gin.Context) {
		wsServer.HandleConnection(c.Writer, c.Request)
	})
	RegisterSensorModule(router, sensorAirePublisher, wsServer)
	RegisterSensorHumedadModule(router, sensorHumedadConsumer, wsServer)
	RegisterNivelAguaModule(router, sensorNivelAguaConsumer, wsServer)
	RegisterSensorUVModule(router, sensorUVConsumer)
	RegisterUserModule(router)
	RegisterSensorTemperaturaModule(router, SensorTemperaturaConsumer, wsServer)
}
