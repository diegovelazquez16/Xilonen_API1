package messaging

import (
	"encoding/json"
	"log"
	"Xilonen-1/sensor/domain/models"
	"Xilonen-1/sensor/aplication/usecase"
	"Xilonen-1/websocket"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorConsumer struct {
	guardarSensorUC *usecase.GuardarSensorUseCase
	wsServer        *websocket.WebSocketServer
	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorConsumer(guardarSensorUC *usecase.GuardarSensorUseCase, wsServer *websocket.WebSocketServer) (*SensorConsumer, error) {
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
		wsServer:        wsServer,
		conn:            conn,
		channel:         ch,
	}, nil
}

func (c *SensorConsumer) Start() {
	msgs, err := c.channel.Consume("aire.procesado", "", true, false, false, false, nil)
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

			err := c.guardarSensorUC.GuardarDatosSensor(sensorData.Valor, sensorData.Categoria, sensorData.Tipo)
			if err != nil {
				log.Printf("❌ Error al guardar en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado: ID=%d, Valor=%.2f", sensorData.ID, sensorData.Valor)
				message := map[string]interface{}{
					"id":         sensorData.ID,
					"valor":      sensorData.Valor,
					"categoria":  sensorData.Categoria,
					"fecha_hora": sensorData.FechaHora,
					"tipo":       "MQ135", 
				}
				c.wsServer.BroadcastMessage("MQ135", message)
			}
		}
	}()
}

func (c *SensorConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
//ok?