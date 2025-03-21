package usecase

import (
	"holamundo/sensor/domain/models"
	"holamundo/sensor/domain/repository"
	"time"
)

type GuardarSensorUseCase struct {
	SensorRepo repository.ISensorRepository
}

func (uc *GuardarSensorUseCase) GuardarDatosSensor(co2, nh3, alcohol, tolueno, acetona float64) error {
	sensor := models.SensorMQ135{
		CO2:       co2,
		NH3:       nh3,
		Alcohol:   alcohol,
		Tolueno:   tolueno,
		Acetona:   acetona,
		FechaHora: time.Now(),
	}

	return uc.SensorRepo.Guardar(&sensor)
}


