package models

import "time"

type SensorMQ135 struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Valor     float64   `json:"valor"`       
	FechaHora time.Time `json:"fecha_hora"` 
}
