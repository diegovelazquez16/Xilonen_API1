package messaging

import (
	"encoding/json"
	"log"
	"os"

	"Xilonen-1/nivelAgua/aplication/usecase"
	"Xilonen-1/nivelAgua/domain/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorNivelAguaConsumer struct {
	guardarSensorUC *usecase.GuardarNivelAguaUseCase
	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorNivelAguaConsumer(guardarSensorUC *usecase.GuardarNivelAguaUseCase) (*SensorNivelAguaConsumer, error) {

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

	return &SensorNivelAguaConsumer{
		guardarSensorUC: guardarSensorUC,
		conn:            conn,
		channel:         ch,
	}, nil
}

//NO olviadar que aqui se inicializa el consumidor y escucha mensajes de la cola "nivelagua.procesado"
func (c *SensorNivelAguaConsumer) Start() {
	msgs, err := c.channel.Consume(
		"nivelagua.procesado", "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes: %v", err)
	}

	go func() {
		for msg := range msgs {
			var sensorData models.SensorT1592
			if err := json.Unmarshal(msg.Body, &sensorData); err != nil {
				log.Printf("⚠️ Error al deserializar el mensaje: %v", err)
				continue
			}

			err := c.guardarSensorUC.GuardarDatosNivelAgua(sensorData.NivelAgua, sensorData.Categoria)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado en BD: ID=%d, Valor=%.2f, FechaHora=%s", sensorData.ID, sensorData.NivelAgua, sensorData.FechaHora)
			}
		}
	}()

	log.Println("📡 Esperando datos de la cola 'nivelagua.procesado'...")
}

func (c *SensorNivelAguaConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
