package publisher

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewEventPublisher(host string) (*EventPublisher, error) {
	conn, err := amqp.Dial(host)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}
	return &EventPublisher{conn: conn, ch: ch}, nil
}

func (p *EventPublisher) Close() {
	p.ch.Close()
	p.conn.Close()
}

func (p *EventPublisher) PublishProductCreated(evt ProductCreatedEvent) error {
	body, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return p.ch.Publish(
		"product.exchange", // exchange
		"",                 // routing key (fanout)
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}