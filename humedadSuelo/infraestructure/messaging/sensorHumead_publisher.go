package messaging

import (
	"encoding/json"
	"log"

	"Xilonen-1/humedadSuelo/aplication/usecase"
	"Xilonen-1/humedadSuelo/domain/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorHumedadConsumer struct {
	guardarSensorUC *usecase.GuardarSensorHumedadUseCase
	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorHumedadConsumer(guardarSensorUC *usecase.GuardarSensorHumedadUseCase) (*SensorHumedadConsumer, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &SensorHumedadConsumer{
		guardarSensorUC: guardarSensorUC,
		conn:            conn,
		channel:         ch,
	}, nil
}

func (c *SensorHumedadConsumer) Start() {
	msgs, err := c.channel.Consume(
		"humedad.procesado", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	go func() {
		for msg := range msgs {
			var sensorData models.SensorLM393
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			// Guardar el dato procesado en la BD usando el caso de uso
			err := c.guardarSensorUC.GuardarDatosSensorHumedad( sensorData.ValorHumedad, sensorData.Categoria)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado en BD: ID=%d, Valor=%.2f, %FechaHora=%s", sensorData.ID, sensorData.ValorHumedad, sensorData.FechaHora)
			}
		}
	}()

	log.Println("📡 Esperando datos de la cola 'humedad.procesado'...")
}

// Close cierra la conexión y el canal de RabbitMQ
func (c *SensorHumedadConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
