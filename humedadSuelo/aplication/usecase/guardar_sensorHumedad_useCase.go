package usecase

import (
	"Xilonen-1/humedadSuelo/domain/models"
	"Xilonen-1/humedadSuelo/domain/repository"
	"time"
)

type GuardarSensorHumedadUseCase struct {
	SensorRepo repository.ISensorHumedadRepository
}

func (uc *GuardarSensorHumedadUseCase) GuardarDatosSensorHumedad(id uint, valor float64, categoria string) error {
	sensorHumedad := models.SensorLM393{
		ID: id,
		ValorHumedad:       valor,
		Categoria: categoria,
		FechaHora: time.Now(),
	}

	return uc.SensorRepo.Guardar(&sensorHumedad)
}


