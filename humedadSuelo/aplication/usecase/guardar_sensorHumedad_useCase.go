package usecase

import (
	"Xilonen-1/humedadSuelo/domain/models"
	"Xilonen-1/humedadSuelo/domain/repository"
	"time"
)

type GuardarSensorHumedadUseCase struct {
	SensorRepo repository.ISensorHumedadRepository
}

func (uc *GuardarSensorHumedadUseCase) GuardarDatosSensorHumedad(valor float64, categoria string) error {
	sensorHumedad := models.SensorLM393{
		ValorHumedad:       valor,
		Categoria: categoria,
		FechaHora: time.Now(),
	}

	return uc.SensorRepo.Guardar(&sensorHumedad)
}


