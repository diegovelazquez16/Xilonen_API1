package repository

import (
	"Xilonen-1/sensorUV/domain/models"
	"gorm.io/gorm"
)



type SensorUVRepositoryImpl struct {
	DB *gorm.DB
}

func (r *SensorUVRepositoryImpl) Guardar(sensorUV *models.SensorUV) error {
	return r.DB.Create(sensorUV).Error
}

func (r *SensorUVRepositoryImpl) ObtenerTodos() ([]models.SensorUV, error) {
	var sensoresUV []models.SensorUV
	err := r.DB.Find(&sensoresUV).Error
	return sensoresUV, err
}
