package main

import (
	"log"
	"Xilonen-1/core"
	"Xilonen-1/launch"

	"Xilonen-1/sensor/infraestructure/messaging"
	sensorHumedadMessaging "Xilonen-1/humedadSuelo/infraestructure/messaging"
	sensorNivelAguaMessaging "Xilonen-1/nivelAgua/infraestructure/messaging"
	sensorUVMessaging "Xilonen-1/sensorUV/infraestructure/messaging"
	sensorTemperaturaMessaging "Xilonen-1/sensorTemperatura/infraestructure/messaging"
	"Xilonen-1/sensor/infraestructure/websocket"

	




	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitializeApp()

	// 🆕 Inicializar WebSocketServer
	wsServer := websocket.NewWebSocketServer()

	// Inicializar el consumidor de aire
	sensorAireConsumer, err := messaging.NewSensorConsumer(nil,wsServer)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Aire: %v", err)
	}

	sensorHumedadConsumer, err := sensorHumedadMessaging.NewSensorHumedadConsumer(nil)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Humedad: %v", err)
	}

	sensorNivelAguaConsumer, err := sensorNivelAguaMessaging.NewSensorNivelAguaConsumer(nil)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Humedad: %v", err)
	}
	sensorUVConsumer, err := sensorUVMessaging.NewSensorUVConsumer(nil)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Humedad: %v", err)
	}
	sensorTemperaturaConsumer, err := sensorTemperaturaMessaging.NewSensorTemperaturaConsumer(nil)
	if err != nil {
		log.Fatalf("❌ Error al conectar con RabbitMQ para Sensor Temperatura: %v", err)
	}



	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Upgrade"}, // 🆕 Permitir Upgrade para WebSocket

	}))

	launch.RegisterRoutes(app, sensorAireConsumer, sensorHumedadConsumer, sensorNivelAguaConsumer, sensorUVConsumer,sensorTemperaturaConsumer, wsServer)

	log.Println("🚀 API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("❌ Error al correr el servidor: %v", err)
	}
}
