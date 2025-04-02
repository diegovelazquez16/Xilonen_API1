package usecase

import (
	"errors"
	"Xilonen-1/nivelAgua/domain/models"
	"Xilonen-1/nivelAgua/domain/repository"
	"time"
)

type GuardarNivelAguaUseCase struct {
	NivelAguaRepo repository.INivelAguaRepository
}

func (uc *GuardarNivelAguaUseCase) GuardarDatosNivelAgua(valorNivelAgua float64, categoria string) error {
	if uc.NivelAguaRepo == nil {
		return errors.New("❌ Error: SensorRepo no ha sido inicializado")
	}
	nivelAgua := models.SensorT1592{
		Categoria: categoria,
		NivelAgua:       valorNivelAgua,
		FechaHora: time.Now(),
	}

	return uc.NivelAguaRepo.Guardar(&nivelAgua)
}


