package repository

import (
	"holamundo/sensor/domain/models"
	"gorm.io/gorm"
)



type SensorRepositoryImpl struct {
	DB *gorm.DB
}

func (r *SensorRepositoryImpl) Guardar(sensor *models.SensorMQ135) error {
	return r.DB.Create(sensor).Error
}

func (r *SensorRepositoryImpl) ObtenerTodos() ([]models.SensorMQ135, error) {
	var sensores []models.SensorMQ135
	err := r.DB.Find(&sensores).Error
	return sensores, err
}
