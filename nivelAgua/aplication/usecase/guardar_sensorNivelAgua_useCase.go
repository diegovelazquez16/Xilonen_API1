package usecase

import (
	"Xilonen-1/nivelAgua/domain/models"
	"Xilonen-1/nivelAgua/domain/repository"
	"time"
)

type GuardarNivelAguaUseCase struct {
	NivelAguaRepo repository.INivelAguaRepository
}

func (uc *GuardarNivelAguaUseCase) GuardarDatosNivelAgua(valorNivelAgua float64) error {
	nivelAgua := models.SensorT1592{
		NivelAgua:       valorNivelAgua,
		FechaHora: time.Now(),
	}

	return uc.NivelAguaRepo.Guardar(&nivelAgua)
}


