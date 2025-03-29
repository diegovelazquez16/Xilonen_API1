package repository

import (
	"Xilonen-1/humedadSuelo/domain/models"

)

type ISensorHumedadRepository interface {
	Guardar(sensor *models.SensorLM393) error
	ObtenerTodos() ([]models.SensorLM393, error)
}




