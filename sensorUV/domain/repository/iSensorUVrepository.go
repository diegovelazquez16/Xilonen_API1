package repository

import (
	"Xilonen-1/sensorUV/domain/models"

)

type ISensorUVRepository interface {
	Guardar(sensor *models.SensorUV) error
	ObtenerTodos() ([]models.SensorUV, error)
}




