package models

import "time"

type SensorLM393 struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ValorHumedad     float64   `json:"valor_humedad"`       
	FechaHora time.Time `json:"fecha_hora"` 
}
