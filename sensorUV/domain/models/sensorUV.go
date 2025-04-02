package models

import "time"

type SensorUV struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ValorUV    float64   `json:"valor_uv"`
	Categoria string  `json:"categoria"`              
	FechaHora time.Time `json:"fecha_hora"` 
}
