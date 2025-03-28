package repository

import (
	"Xilonen-1/nivelAgua/domain/models"

)

type INivelAguaRepository interface {
	Guardar(sensor *models.SensorT1592) error
	ObtenerTodos() ([]models.SensorT1592, error)
}




