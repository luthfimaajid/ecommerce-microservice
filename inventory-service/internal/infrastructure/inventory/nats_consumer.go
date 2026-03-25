package inventory

import (
	"context"
	"ecom-ms/inventory-service/internal/domain/inventory"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type OrderCreatedHandler interface {
	HandleOrderCreated(ctx context.Context, event inventory.OrderCreatedEvent) error
}

type NatsConsumer struct {
	conn    *nats.Conn
	handler OrderCreatedHandler
}

func NewNatsConsumer(conn *nats.Conn, handler OrderCreatedHandler) *NatsConsumer {
	return &NatsConsumer{
		conn:    conn,
		handler: handler,
	}
}

func (c *NatsConsumer) Start() error {

	_, err := c.conn.Subscribe("order.created", func(msg *nats.Msg) {

		var event inventory.OrderCreatedEvent

		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Println("failed to decode event:", err)
			return
		}

		err = c.handler.HandleOrderCreated(context.Background(), event)
		if err != nil {
			log.Println("handler error:", err)
		}

	})

	return err
}
