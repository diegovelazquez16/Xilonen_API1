package models

import "time"

type SensorMQ135 struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CO2       float64   `json:"co2"`       // Concentración de CO2 en ppm
	NH3       float64   `json:"nh3"`       // Concentración de Amoníaco en ppm
	Alcohol   float64   `json:"alcohol"`   // Concentración de Alcohol en ppm
	Tolueno   float64   `json:"tolueno"`   // Concentración de Tolueno en ppm
	Acetona   float64   `json:"acetona"`   // Concentración de Acetona en ppm
	FechaHora time.Time `json:"fecha_hora"` // Timestamp de la medición
}
