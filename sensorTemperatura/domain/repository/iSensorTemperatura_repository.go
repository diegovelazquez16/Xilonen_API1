package repository

import (
	"Xilonen-1/sensorTemperatura/domain/models"

)

type ISensorTemperaturaRepository interface {
	Guardar(sensor *models.SensorDHT11) error
	ObtenerTodos() ([]models.SensorDHT11, error)
}




