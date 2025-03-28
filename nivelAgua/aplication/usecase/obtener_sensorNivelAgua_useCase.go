package usecase

import (
	"Xilonen-1/nivelAgua/domain/models"
	"Xilonen-1/nivelAgua/domain/repository"
)


type ObtenerNivelAguaUseCase struct {
	NivelAguaRepo repository.INivelAguaRepository
}
func (uc *ObtenerNivelAguaUseCase) ObtenerNivelAgua() ([]models.SensorT1592, error) {
	return uc.NivelAguaRepo.ObtenerTodos()
}
