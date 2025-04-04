package usecase

import (
	"Xilonen-1/sensorTemperatura/domain/models"
	"Xilonen-1/sensorTemperatura/domain/repository"
)


type ObtenerSensorTemperaturaUseCase struct {
	SensorRepo repository.ISensorTemperaturaRepository
}
func (uc *ObtenerSensorTemperaturaUseCase) ObtenerDatosSensoresTemperatura() ([]models.SensorDHT11, error) { 
	return uc.SensorRepo.ObtenerTodos()
}
