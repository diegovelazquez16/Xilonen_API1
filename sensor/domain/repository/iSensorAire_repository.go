package repository

import (
	"holamundo/sensor/domain/models"

)

type ISensorRepository interface {
	Guardar(sensor *models.SensorMQ135) error
	ObtenerTodos() ([]models.SensorMQ135, error)
}




