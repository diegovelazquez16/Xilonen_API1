package usecase

import (
	"Xilonen-1/sensorTemperatura/domain/models"
	"Xilonen-1/sensorTemperatura/domain/repository"
	"time"
)

type GuardarSensorTemperaturaUseCase struct {
	SensorRepo repository.ISensorTemperaturaRepository
}

func (uc *GuardarSensorTemperaturaUseCase) GuardarDatosSensorTemperatura(id uint, valor float64, categoria string) error {
	sensorTemperatura := models.SensorDHT11{
		ID: id,
		ValorTemperatura:       valor,
		Categoria: categoria,
		FechaHora: time.Now(),
	}

	return uc.SensorRepo.Guardar(&sensorTemperatura)
}


