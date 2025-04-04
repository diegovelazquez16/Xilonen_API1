package simulator

import (
	"log"
	"math/rand"
	"time"

	"Xilonen-1/sensor/infraestructure/websocket"
)

// SimulateSensorData envía datos simulados de sensores al WebSocket
func SimulateSensorData(wsServer *websocket.WebSocketServer) {
	sensorTypes := []string{"MQ135", "LM395", "T1592", "DHT11"}
	categories := []string{"Bueno", "Moderado", "Peligroso"}

	for {
		// Selecciona un tipo de sensor aleatorio
		sensorType := sensorTypes[rand.Intn(len(sensorTypes))]

		// Genera datos simulados
		data := map[string]interface{}{
			"id":         rand.Intn(1000),          // ID aleatorio
			"valor":      float64(rand.Intn(5000)), // Valor del sensor aleatorio
			"categoria":  categories[rand.Intn(len(categories))], // Categoría aleatoria
			"fecha_hora": time.Now().Format(time.RFC3339), // Fecha y hora en formato estándar
			"tipo":       sensorType, // Tipo de sensor
		}

		log.Printf("📡 Simulando datos: Sensor=%s, ID=%d, Valor=%.2f, Categoría=%s, Fecha=%s",
			sensorType, data["id"], data["valor"], data["categoria"], data["fecha_hora"])

		// Enviar datos al WebSocket
		wsServer.BroadcastMessage(sensorType, data)

		// Esperar 3 segundos antes de enviar otro dato
		time.Sleep(3 * time.Second)
	}
}
