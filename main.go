package main

import (
	"log"
	"Xilonen-1/core"
	"Xilonen-1/launch"

	"Xilonen-1/sensor/infraestructure/messaging"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"


	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitializeApp()

	// Inicializar el consumidor de aire
	sensorAireConsumer, err := messaging.NewSensorConsumer(nil)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Aire: %v", err)
	}

	sensorHumedadConsumer, err := sensorHumedadMessaging.NewSensorHumedadConsumer(nil)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Humedad: %v", err)
	}







	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	launch.RegisterRoutes(app, sensorAireConsumer, sensorHumedadConsumer)

	log.Println("🚀 API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("❌ Error al correr el servidor: %v", err)
	}
}
