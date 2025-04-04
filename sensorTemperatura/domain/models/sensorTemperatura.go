package models

import "time"

type SensorDHT11 struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ValorTemperatura     float64   `json:"valor_temperatura"`
	Categoria string  `json:"categoria"`       
	FechaHora time.Time `json:"fecha_hora"` 
}
