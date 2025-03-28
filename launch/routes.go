// launch/routes.go
package launch

import (
	"github.com/gin-gonic/gin"
	"Xilonen-1/sensor/infraestructure/messaging"

)



func RegisterRoutes(router *gin.Engine, sensorAirePublisher * messaging.SensorConsumer) {

	RegisterSensorModule(router, sensorAirePublisher)
	RegisterSensorHumedadModule(router)
	RegisterNivelAguaModule(router)
	RegisterSensorUVModule(router)

}
