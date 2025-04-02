package messaging

import (
	"encoding/json"
	"log"
	"os"
	"Xilonen-1/sensor/aplication/usecase"
	"Xilonen-1/sensor/domain/models"
	"Xilonen-1/sensor/infraestructure/websocket"

	amqp "github.com/rabbitmq/amqp091-go"
)

// SensorConsumer maneja la conexión con RabbitMQ
type SensorConsumer struct {
	guardarSensorUC *usecase.GuardarSensorUseCase
	conn            *amqp.Connection
	channel         *amqp.Channel
	wsServer        *websocket.WebSocketServer
}

func NewSensorConsumer(guardarSensorUC *usecase.GuardarSensorUseCase, wsServer *websocket.WebSocketServer) (*SensorConsumer, error) {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("❌ ERROR: RABBITMQ_URL no está configurada")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &SensorConsumer{
		guardarSensorUC: guardarSensorUC,
		conn:            conn,
		channel:         ch,
		wsServer:        wsServer,
	}, nil
}

func (c *SensorConsumer) Start(wsServer *websocket.WebSocketServer) {
	msgs, err := c.channel.Consume(
		"aire.procesado", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	go func() {
		for msg := range msgs {
			var sensorData models.SensorMQ135
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			err := c.guardarSensorUC.GuardarDatosSensor(sensorData.Valor, sensorData.Categoria)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado en BD: ID=%d, Valor=%.2f, FechaHora=%s",
					sensorData.ID, sensorData.Valor, sensorData.FechaHora)

				wsServer.SendSensorData(sensorData)
			}
		}
	}()
	log.Println("📡 Esperando datos de la cola 'aire.procesado'...")
}

//ok