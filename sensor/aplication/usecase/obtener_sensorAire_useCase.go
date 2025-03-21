package usecase

import (
	"holamundo/sensor/domain/models"
	"holamundo/sensor/domain/repository"
)


type ObtenerSensorUseCase struct {
	SensorRepo repository.ISensorRepository
}
func (uc *ObtenerSensorUseCase) ObtenerDatosSensores() ([]models.SensorMQ135, error) {
	return uc.SensorRepo.ObtenerTodos()
}
