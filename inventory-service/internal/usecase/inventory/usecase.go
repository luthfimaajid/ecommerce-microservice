package inventory

import (
	"context"
	"ecom-ms/inventory-service/internal/domain/inventory"
	"time"
)

type OrderEventHandler struct {
	repo inventory.Repository
}

func NewOrderEventHandler(repo inventory.Repository) *OrderEventHandler {
	return &OrderEventHandler{
		repo: repo,
	}
}

func (h *OrderEventHandler) HandleOrderCreated(
	ctx context.Context,
	event inventory.OrderCreatedEvent,
) error {
	tx, err := h.repo.CreateTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stockHistory := inventory.StockHistory{
		ProductID:    event.ProductID,
		OrderID:      &event.OrderID,
		UserID:       event.UserID,
		MovementType: "order_pending",
		Quantity:     -event.Quantity,
		CreatedAt:    time.Now(),
	}

	return h.repo.InsertStockHistory(
		ctx,
		tx,
		&stockHistory,
	)
}
