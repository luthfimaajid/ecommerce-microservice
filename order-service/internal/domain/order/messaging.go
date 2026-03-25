package order

import "context"

type Messaging interface {
	PublishOrderCreated(ctx context.Context, event OrderCreatedEvent) error
}
