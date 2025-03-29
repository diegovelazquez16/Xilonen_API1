package repository

import (
	"Xilonen-1/nivelAgua/domain/models"
	"gorm.io/gorm"
)



type NivelAguaRepositoryImpl struct {
	DB *gorm.DB
}

func (r *NivelAguaRepositoryImpl) Guardar(NivelAgua *models.SensorT1592) error {
	return r.DB.Create(NivelAgua).Error
}

func (r *NivelAguaRepositoryImpl) ObtenerTodos() ([]models.SensorT1592, error) {
	var sensoresNivelesAgua []models.SensorT1592
	err := r.DB.Find(&sensoresNivelesAgua).Error
	return sensoresNivelesAgua, err
}
