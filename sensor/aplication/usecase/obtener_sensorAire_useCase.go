package usecase

import (
	"Xilonen-1/sensor/domain/models"
	"Xilonen-1/sensor/domain/repository"
)


type ObtenerSensorUseCase struct {
	SensorRepo repository.ISensorRepository
}
func (uc *ObtenerSensorUseCase) ObtenerDatosSensores() ([]models.SensorMQ135, error) {
	return uc.SensorRepo.ObtenerTodos()
}
