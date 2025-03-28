package usecase

import (
	"Xilonen-1/humedadSuelo/domain/models"
	"Xilonen-1/humedadSuelo/domain/repository"
)


type ObtenerSensorHumedadUseCase struct {
	SensorRepo repository.ISensorHumedadRepository
}
func (uc *ObtenerSensorHumedadUseCase) ObtenerDatosSensoresHumedad() ([]models.SensorLM393, error) { 
	return uc.SensorRepo.ObtenerTodos()
}
