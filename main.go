package main

import (
	"log"
	"Xilonen-1/core"
	"Xilonen-1/launch"
	"Xilonen-1/sensor/aplication/usecase"
	"Xilonen-1/sensor/domain/repository"
	"Xilonen-1/sensor/infraestructure/messaging"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitializeApp()

	// 🔹 Crear repositorio y caso de uso para sensores
	sensorRepo := &repository.SensorRepositoryImpl{DB: core.GetDB()}
	guardarSensorUC := &usecase.GuardarSensorUseCase{SensorRepo: sensorRepo}



	// 🔹 Inicializar el consumidor de sensores
	sensorConsumer, err := messaging.NewSensorConsumer(guardarSensorUC)
	if err != nil {
		log.Fatalf("❌ Error al iniciar el consumidor de sensores: %v", err)
	}
	defer sensorConsumer.Close()
	
	// Ejecutar el consumidor en un goroutine
	go sensorConsumer.Start()

	// 🔹 Configurar API con Gin
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 🔹 Registrar rutas con el publicador de sensores
	launch.RegisterRoutes(app, sensorConsumer)

	log.Println("🚀 API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("❌ Error al correr el servidor: %v", err)
	}
}
