// launch/routes.go
package launch

import (
	"github.com/gin-gonic/gin"
	"Xilonen-1/sensor/infraestructure/messaging"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"
	sensorNivelAguaMessaging "Xilonen-1/nivelAgua/infraestructure/messaging"

)



func RegisterRoutes(router *gin.Engine, sensorAirePublisher *messaging.SensorConsumer, sensorHumedadConsumer *sensorHumedadMessaging.SensorHumedadConsumer, sensorNivelAguaConsumer *sensorNivelAguaMessaging.SensorNivelAguaConsumer) {

	RegisterSensorModule(router, sensorAirePublisher)
	RegisterSensorHumedadModule(router, sensorHumedadConsumer)
	RegisterNivelAguaModule(router, sensorNivelAguaConsumer)
	RegisterSensorUVModule(router)

}
