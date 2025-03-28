package models

import "time"

type SensorT1592 struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NivelAgua     float64   `json:"nivel_agua"`       
	FechaHora time.Time `json:"fecha_hora"` 
}
