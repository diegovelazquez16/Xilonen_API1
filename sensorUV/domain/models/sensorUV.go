package models

import "time"

type SensorUV struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ValorUV    float64   `json:"valor_uv"`       
	FechaHora time.Time `json:"fecha_hora"` 
}
