package launch

import (
"github.com/gin-gonic/gin"
"holamundo/sensor/infraestructure/messaging"

)



func RegisterRoutes(router *gin.Engine, sensorAirePublisher * messaging.SensorPublisher) {

	RegisterSensorModule(router, sensorAirePublisher)

}
