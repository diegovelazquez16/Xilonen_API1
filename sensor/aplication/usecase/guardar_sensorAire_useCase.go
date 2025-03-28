package usecase

import (
	"Xilonen-1/sensor/domain/models"
	"Xilonen-1/sensor/domain/repository"
	"time"
)

type GuardarSensorUseCase struct {
	SensorRepo repository.ISensorRepository
}

func (uc *GuardarSensorUseCase) GuardarDatosSensor(valor float64, categoria string) error {
	sensor := models.SensorMQ135{
		Valor:       valor,
		Categoria: categoria,
		FechaHora: time.Now(),
	}

	return uc.SensorRepo.Guardar(&sensor)
}


