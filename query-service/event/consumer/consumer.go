package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/NhuqyGit/cqrs-order-demo/query-service/models"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewEventConsumer(host string) (*EventConsumer, error) {
	conn, err := amqp.Dial(host)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Declare exchange
	err = ch.ExchangeDeclare(
		"product.exchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &EventConsumer{conn: conn, ch: ch}, nil
}

func (c *EventConsumer) Close() {
	c.ch.Close()
	c.conn.Close()
}

func (c *EventConsumer) StartProductCreatedConsumer(productService service.ProductService) error {
	q, err := c.ch.QueueDeclare(
		"", false, true, false, false, nil,
	)
	if err != nil {
		return err
	}

	err = c.ch.QueueBind(
		q.Name,
		"", // fanout → routing key ignored
		"product.exchange",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := c.ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			var evt ProductCreatedEvent
			if err := json.Unmarshal(d.Body, &evt); err != nil {
				log.Println("Error decoding product event:", err)
				continue
			}

			product := serviceToModel(evt)
			err := productService.CreateProduct(context.Background(), product)
			if err != nil {
				log.Println("Failed to save product to MongoDB:", err)
			}
			log.Println("Create product to MongoDB success")
		}
	}()

	return nil
}

// Mapping event to model
func serviceToModel(evt ProductCreatedEvent) models.Product {
	return models.Product{
		ProductID:      evt.ID,
		Name:     		evt.Name,
		Description: 	evt.Description,
		Price:    		evt.Price,
		Quantity: 		evt.Quantity,
		SKU: 			evt.SKU,
	}
}
