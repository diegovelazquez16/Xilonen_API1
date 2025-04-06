package messaging

import (
	"encoding/json"
	"log"

	"Xilonen-1/sensorTemperatura/aplication/usecase"
	"Xilonen-1/sensorTemperatura/domain/models"
	"Xilonen-1/sensor/infraestructure/websocket"


	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorTemperaturaConsumer struct {
	guardarSensorUC *usecase.GuardarSensorTemperaturaUseCase
	wsServer 		* websocket.WebSocketServer

	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorTemperaturaConsumer(guardarSensorUC *usecase.GuardarSensorTemperaturaUseCase, wsServer * websocket.WebSocketServer) (*SensorTemperaturaConsumer, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &SensorTemperaturaConsumer{
		guardarSensorUC: guardarSensorUC,
		wsServer: wsServer,
		conn:            conn,
		channel:         ch,
	}, nil
}

func (c *SensorTemperaturaConsumer) Start() {
	msgs, err := c.channel.Consume(
		"temperatura.procesado", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	go func() {
		for msg := range msgs {
			var sensorData models.SensorDHT11
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			// Guardar el dato procesado en la BD usando el caso de uso
			err := c.guardarSensorUC.GuardarDatosSensorTemperatura( sensorData.ID, sensorData.ValorTemperatura, sensorData.Categoria, sensorData.Tipo)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado: ID=%d, Valor=%.2f", sensorData.ID, sensorData.ValorTemperatura)
				message := map[string]interface{}{
					"id":         sensorData.ID,
					"valor":      sensorData.ValorTemperatura,
					"categoria":  sensorData.Categoria,
					"fecha_hora": sensorData.FechaHora,
					"tipo":       "Humedad", // Identifica el sensor de calidad de aire
				}
				// Enviar al WebSocket
				c.wsServer.BroadcastMessage("Humedad", message)
			}
		}
	}()

	log.Println("📡 Esperando datos de la cola 'temperatura.procesado'...")
}

// Close cierra la conexión y el canal de RabbitMQ
func (c *SensorTemperaturaConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
