package models

import "time"

type SensorT1592 struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NivelAgua     float64   `json:"nivel_agua"`  
	Categoria string  `json:"categoria"`       
	FechaHora time.Time `json:"fecha_hora"` 
	Tipo string `json:"tipo"`
}
