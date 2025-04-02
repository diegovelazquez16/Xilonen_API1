package repository

import (
	"Xilonen-1/sensorTemperatura/domain/models"
	"gorm.io/gorm"
)



type SensorTemperaturaRepositoryImpl struct {
	DB *gorm.DB
}

func (r *SensorTemperaturaRepositoryImpl) Guardar(sensorTemperatura *models.SensorDHT11) error {
	return r.DB.Create(sensorTemperatura).Error
}

func (r *SensorTemperaturaRepositoryImpl) ObtenerTodos() ([]models.SensorDHT11, error) {
	var sensoresTemperatura []models.SensorDHT11
	err := r.DB.Find(&sensoresTemperatura).Error
	return sensoresTemperatura, err
}
