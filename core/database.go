// core/database.go 
package core

import (
	"fmt"
	"log"
	"os"

	sensorModels "Xilonen-1/sensor/domain/models"
	sensorHumedadModels "Xilonen-1/humedadSuelo/domain/models"
	sensorNivelAguaModels "Xilonen-1/nivelAgua/domain/models"
	sensorUVModels "Xilonen-1/sensorUV/domain/models"
	userModels "Xilonen-1/users/domain/models"
	sensorTemperaturaModels "Xilonen-1/sensorTemperatura/domain/models"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func DatabaseConnection() {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectándose a la base de datos: %v", err)
	}

	DB = db
	log.Println("Conexión a la base de datos exitosa.")
	err = DB.AutoMigrate(&sensorModels.SensorMQ135{}, &sensorHumedadModels.SensorLM393{}, &sensorNivelAguaModels.SensorT1592{}, &sensorUVModels.SensorUV{}, &userModels.User{}, &sensorTemperaturaModels.SensorDHT11{})
	if err != nil {
		log.Fatalf("Error aplicando migración: %v", err)
	}
	log.Println("Migración aplicada exitosamente.")
}

func GetDB() *gorm.DB {
	return DB
}
