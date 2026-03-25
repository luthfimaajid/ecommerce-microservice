package inventory

import (
	"ecom-ms/inventory-service/internal/domain/inventory"
	"time"
)

type StockHistoryModel struct {
	ID           int64     `db:"id"`
	ProductID    int64     `db:"product_id"`
	OrderID      *int64    `db:"order_id"`
	UserID       int64     `db:"user_id"`
	MovementType string    `db:"movement_type"`
	Quantity     int       `db:"quantity"`
	CreatedAt    time.Time `db:"created_at"`
}

func (m *StockHistoryModel) ToDomain() inventory.StockHistory {
	return inventory.StockHistory{
		ID:           m.ID,
		ProductID:    m.ProductID,
		OrderID:      m.OrderID,
		UserID:       m.UserID,
		MovementType: m.MovementType,
		Quantity:     m.Quantity,
		CreatedAt:    m.CreatedAt,
	}
}

func FromDomain(s inventory.StockHistory) StockHistoryModel {
	return StockHistoryModel{
		ID:           s.ID,
		ProductID:    s.ProductID,
		OrderID:      s.OrderID,
		UserID:       s.UserID,
		MovementType: s.MovementType,
		Quantity:     s.Quantity,
		CreatedAt:    s.CreatedAt,
	}
}
