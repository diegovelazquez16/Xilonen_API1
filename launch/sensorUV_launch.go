package launch

import (
	"Xilonen-1/core"

//sensor de UV:
	sensorUVUsecase "Xilonen-1/sensorUV/aplication/usecase"
	sensorUVRepo "Xilonen-1/sensorUV/domain/repository"
	sensorUVControllers "Xilonen-1/sensorUV/infraestructure/controllers"
	sensorUVRoutes "Xilonen-1/sensorUV/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterSensorUVModule(router *gin.Engine) {
	sensorUVRepo := &sensorUVRepo.SensorUVRepositoryImpl{DB: core.GetDB()}

	guardarSensorUVUC := &sensorUVUsecase.GuardarSensorUVUseCase{SensorUVRepo: sensorUVRepo}
	obtenerSensoresUVUC := &sensorUVUsecase.ObtenerSensorUVUseCase{SensorUVRepo: sensorUVRepo}

	guardarSensorUVController := &sensorUVControllers.GuardarSensorUVController{GuardarSensorUVUC: guardarSensorUVUC}
	obtenerSensoresUVController := &sensorUVControllers.ObtenerSensorUVController{ObtenerSensorUVUC: obtenerSensoresUVUC}

	sensorUVRoutes.SensorUVRoutes(router, guardarSensorUVController,obtenerSensoresUVController )
}
//UV