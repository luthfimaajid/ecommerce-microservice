package order

import (
	"context"
	"ecom-ms/order-service/internal/domain/order"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	conn *nats.Conn
}

func NewNatsPublisher(conn *nats.Conn) *NatsPublisher {
	return &NatsPublisher{
		conn: conn,
	}
}

func (p *NatsPublisher) PublishOrderCreated(
	ctx context.Context,
	event order.OrderCreatedEvent,
) error {

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	subject := "order.created"

	return p.conn.Publish(subject, data)
}
