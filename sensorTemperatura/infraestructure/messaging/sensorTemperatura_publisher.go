package messaging

import (
	"encoding/json"
	"log"

	"Xilonen-1/sensorTemperatura/aplication/usecase"
	"Xilonen-1/sensorTemperatura/domain/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorTemperaturaConsumer struct {
	guardarSensorUC *usecase.GuardarSensorTemperaturaUseCase
	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorTemperaturaConsumer(guardarSensorUC *usecase.GuardarSensorTemperaturaUseCase) (*SensorTemperaturaConsumer, error) {
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
			err := c.guardarSensorUC.GuardarDatosSensorTemperatura( sensorData.ID, sensorData.ValorTemperatura, sensorData.Categoria)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado en BD: ID=%d, Valor=%.2f, FechaHora=%s",
					sensorData.ID, sensorData.ValorTemperatura, sensorData.FechaHora)
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
