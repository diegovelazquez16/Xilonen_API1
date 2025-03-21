package messaging

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"holamundo/sensor/domain/models"
)

type SensorPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewSensorPublisher() (*SensorPublisher, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"sensores",  
		true,       
		false,      
		false,      
		false,      
		nil,        
	)
	if err != nil {
		return nil, err
	}

	return &SensorPublisher{conn: conn, channel: ch, queue: q}, nil
}

func (sp *SensorPublisher) Publish(sensor models.SensorMQ135) error {
	body, err := json.Marshal(sensor)
	if err != nil {
		return err
	}

	err = sp.channel.Publish(
		"",           
		sp.queue.Name, 
		false,       
		false,       
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf("Datos enviados a RabbitMQ: %+v", sensor)
	return nil
}

func (sp *SensorPublisher) Close() {
	sp.channel.Close()
	sp.conn.Close()
}
