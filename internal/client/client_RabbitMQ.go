package client

import (
	someRestApi "SomeRestApi"
	"SomeRestApi/internal/config"
	"SomeRestApi/internal/service"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Client struct {
	service *service.Service
}

func NewClient(service *service.Service) *Client {
	return &Client{service: service}
}

func (c *Client) RunClient(cfg config.RabbitMQ) (string, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return "Failed to connect to RabbitMQ", err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return "Failed to open a RabbitMQ channel", err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"create", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return "Failed to declare a RabbitMQ queue", err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return "Failed to register a RabbitMQ consumer", err
	}

	var contact someRestApi.Contacts

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &contact)
			if err != nil {
				log.Panicf("Failed to unmarshal json: %s", err)
			}
			id, err := c.service.CreateContact(contact)
			if err != nil {
				log.Panicf("Service error: %s", err)
			}
			log.Printf("Contact created with id: %d", id)
		}
	}()

	<-forever
	return "", err
}
