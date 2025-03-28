// launch/sensorNivelAgua_launch.go
package launch

import (
	"Xilonen-1/core"

//sensor de nivel de agua:
	nivelAguaUsecase "Xilonen-1/nivelAgua/aplication/usecase"
	nivelAguaRepo "Xilonen-1/nivelAgua/domain/repository"
	nivelAguaControllers "Xilonen-1/nivelAgua/infraestructure/controllers"
	nivelAguaRoutes "Xilonen-1/nivelAgua/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterNivelAguaModule(router *gin.Engine) {
	nivelAguaRepo := &nivelAguaRepo.NivelAguaRepositoryImpl{DB: core.GetDB()}

	guardarNivelAguaUC := &nivelAguaUsecase.GuardarNivelAguaUseCase{NivelAguaRepo: nivelAguaRepo}
	obtenerNivelAguaUC := &nivelAguaUsecase.ObtenerNivelAguaUseCase{NivelAguaRepo: nivelAguaRepo}

	guardarNivelAguaController := &nivelAguaControllers.GuardarNivelAguaController{GuardarNivelAguaUC: guardarNivelAguaUC}
	obtenerNivelAguaController := &nivelAguaControllers.ObtenerNivelAguaController{ObtenerNivelAguaUC: obtenerNivelAguaUC}

	nivelAguaRoutes.NivelAguaRoutes(router, guardarNivelAguaController,obtenerNivelAguaController )
}
