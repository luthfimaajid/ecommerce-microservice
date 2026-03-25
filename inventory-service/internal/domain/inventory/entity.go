package inventory

import "time"

type StockHistory struct {
	ID           int64
	ProductID    int64
	OrderID      *int64
	UserID       int64
	MovementType string
	Quantity     int
	CreatedAt    time.Time
}

type OrderCreatedEvent struct {
	UserID    int64 `json:"user_id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
