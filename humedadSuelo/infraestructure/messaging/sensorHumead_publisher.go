package messaging

import (
	"encoding/json"
	"log"

	"Xilonen-1/humedadSuelo/aplication/usecase"
	"Xilonen-1/humedadSuelo/domain/models"
	"Xilonen-1/websocket"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SensorHumedadConsumer struct {
	guardarSensorUC *usecase.GuardarSensorHumedadUseCase
	wsServer 		* websocket.WebSocketServer
	conn            *amqp.Connection
	channel         *amqp.Channel
}

func NewSensorHumedadConsumer(guardarSensorUC *usecase.GuardarSensorHumedadUseCase, wsServer * websocket.WebSocketServer) (*SensorHumedadConsumer, error) {
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
		wsServer: wsServer,
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

			err := c.guardarSensorUC.GuardarDatosSensorHumedad( sensorData.ID, sensorData.ValorHumedad, sensorData.Categoria, sensorData.Tipo)
			if err != nil {
				log.Printf("❌ Error al guardar el dato en la BD: %v", err)
			} else {
				log.Printf("✅ Dato guardado (humedad): ID=%d, Valor=%.2f", sensorData.ID, sensorData.ValorHumedad)
				message := map[string]interface{}{
					"id":         sensorData.ID,
					"valor":      sensorData.ValorHumedad,
					"categoria":  sensorData.Categoria,
					"fecha_hora": sensorData.FechaHora,
					"tipo":       "Humedad", 
				}
				c.wsServer.BroadcastMessage("Humedad", message)
			}
		}
	}()

	log.Println("📡 Esperando datos de la cola 'humedad.procesado'...")
}

func (c *SensorHumedadConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
