package messaging

import (
	"encoding/json"
	"log"

	"Xilonen-1/sensor/aplication/usecase"
	"Xilonen-1/sensor/domain/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

// SensorConsumer estructura para manejar la conexión con RabbitMQ
type SensorConsumer struct {
	guardarSensorUC *usecase.GuardarSensorUseCase
	conn            *amqp.Connection
	channel         *amqp.Channel
}

// NewSensorConsumer crea un nuevo consumidor de la cola "aire.procesado"
func NewSensorConsumer(guardarSensorUC *usecase.GuardarSensorUseCase) (*SensorConsumer, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
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
	}, nil
}

// Start inicia el consumidor y escucha mensajes de la cola "aire.procesado"
func (c *SensorConsumer) Start() {
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

			// Guardar el dato procesado en la BD usando el caso de uso
			err := c.guardarSensorUC.GuardarDatosSensor(sensorData.Valor, sensorData.Categoria)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado en BD: ID=%d, Valor=%.2f, FechaHora=%s", sensorData.ID, sensorData.Valor, sensorData.FechaHora)
			}
		}
	}()

	log.Println("📡 Esperando datos de la cola 'aire.procesado'...")
}

// Close cierra la conexión y el canal de RabbitMQ
func (c *SensorConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
