package usecase

import (
	"Xilonen-1/sensorUV/domain/models"
	"Xilonen-1/sensorUV/domain/repository"
	"time"
)

type GuardarSensorUVUseCase struct {
	SensorUVRepo repository.ISensorUVRepository
}

func (uc *GuardarSensorUVUseCase) GuardarDatosSensorUV(valorUV float64) error {
	sensorUV := models.SensorUV{
		ValorUV:       valorUV,
		FechaHora: time.Now(),
	}

	return uc.SensorUVRepo.Guardar(&sensorUV)
}


