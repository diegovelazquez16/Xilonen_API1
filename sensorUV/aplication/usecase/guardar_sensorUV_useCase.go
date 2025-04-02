package usecase

import (
	"errors"
	"Xilonen-1/sensorUV/domain/models"
	"Xilonen-1/sensorUV/domain/repository"
	"time"
)

type GuardarSensorUVUseCase struct {
	SensorUVRepo repository.ISensorUVRepository
}

func (uc *GuardarSensorUVUseCase) GuardarDatosSensorUV(valorUV float64, categoria string) error {

	if uc.SensorUVRepo == nil {
		return errors.New("❌ Error: SensorRepo no ha sido inicializado")
	}
	sensorUV := models.SensorUV{
		Categoria: categoria,
		ValorUV:       valorUV,
		FechaHora: time.Now(),
	}
	return uc.SensorUVRepo.Guardar(&sensorUV)
}


