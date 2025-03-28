// launch/routes.go
package launch

import (
	"github.com/gin-gonic/gin"
	"Xilonen-1/sensor/infraestructure/messaging"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"

)



func RegisterRoutes(router *gin.Engine, sensorAirePublisher *messaging.SensorConsumer, sensorHumedadConsumer *sensorHumedadMessaging.SensorHumedadConsumer) {

	RegisterSensorModule(router, sensorAirePublisher)
	RegisterSensorHumedadModule(router, sensorHumedadConsumer)
	RegisterNivelAguaModule(router)
	RegisterSensorUVModule(router)

}
