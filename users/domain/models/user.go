package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
    Nombre    string `gorm:"not null"`
    Email     string `gorm:"unique;not null"`
    Telefono  string
    Direccion string
    Password string  `gorm:"unique;not null"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}