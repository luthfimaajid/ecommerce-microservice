package order

import "time"

type Product struct {
	ID          int64
	Category    string
	Name        string
	Description *string
	Price       float64
	Stock       int
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Status string

const (
	StatusPending   Status = "pending"
	StatusSuccess   Status = "success"
	StatusCancelled Status = "cancelled"
)

type Order struct {
	ID          int64
	UserID      int64
	ProductID   int64
	Quantity    int
	UnitPrice   float64
	TotalAmount float64
	Status      string
	Note        *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderCreatedEvent struct {
	UserID    int64 `json:"user_id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
