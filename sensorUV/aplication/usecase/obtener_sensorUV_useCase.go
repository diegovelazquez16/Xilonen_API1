package usecase

import (
	"Xilonen-1/sensorUV/domain/models"
	"Xilonen-1/sensorUV/domain/repository"
)


type ObtenerSensorUVUseCase struct {
	SensorUVRepo repository.ISensorUVRepository
}
func (uc *ObtenerSensorUVUseCase) ObtenerDatosSensoresUV() ([]models.SensorUV, error) { 
	return uc.SensorUVRepo.ObtenerTodos()
}
