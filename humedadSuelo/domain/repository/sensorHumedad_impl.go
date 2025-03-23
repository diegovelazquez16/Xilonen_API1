package repository

import (
	"Xilonen-1/humedadSuelo/domain/models"
	"gorm.io/gorm"
)



type SensorHumedadRepositoryImpl struct {
	DB *gorm.DB
}

func (r *SensorHumedadRepositoryImpl) Guardar(sensorHumedad *models.SensorLM393) error {
	return r.DB.Create(sensorHumedad).Error
}

func (r *SensorHumedadRepositoryImpl) ObtenerTodos() ([]models.SensorLM393, error) {
	var sensoresHumedad []models.SensorLM393
	err := r.DB.Find(&sensoresHumedad).Error
	return sensoresHumedad, err
}
