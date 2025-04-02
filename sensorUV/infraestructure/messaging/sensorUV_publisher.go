package messaging

import (
	"encoding/json"
	"log"
	"os"

	"Xilonen-1/sensorUV/aplication/usecase"
	"Xilonen-1/sensorUV/domain/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorUVConsumer struct {
	guardarSensorUC *usecase.GuardarSensorUVUseCase
	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorUVConsumer(guardarSensorUC *usecase.GuardarSensorUVUseCase) (*SensorUVConsumer, error) {

	rabbitURL := os.Getenv("RABBITMQ_URL") //NO olvidar cargar variables de entorno
	if rabbitURL == "" {
		log.Fatal("❌ ERROR: RABBITMQ_URL no está configurada en el entorno")
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

	return &SensorUVConsumer{
		guardarSensorUC: guardarSensorUC,
		conn:            conn,
		channel:         ch,
	}, nil
}

//NO olviadar que aqui se inicializa el consumidor y escucha mensajes de la cola "luz.procesado"
func (c *SensorUVConsumer) Start() {
	msgs, err := c.channel.Consume(
		"uv.procesado", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	go func() {
		for msg := range msgs {
			var sensorData models.SensorUV
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			err := c.guardarSensorUC.GuardarDatosSensorUV(sensorData.ValorUV, sensorData.Categoria)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado en BD: ID=%d, Valor=%.2f, FechaHora=%s", sensorData.ID, sensorData.ValorUV, sensorData.FechaHora)
			}
		}
	}()

	log.Println("📡 Esperando datos de la cola 'luzuv.procesado'...")
}

func (c *SensorUVConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
