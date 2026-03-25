package order

import (
	orderDom "ecom-ms/order-service/internal/domain/order"
	"time"
)

type Product struct {
	ID          int64     `db:"id"`
	Category    string    `db:"category"`
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	IsActive    bool      `db:"is_active"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Order struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	ProductID   int64     `db:"product_id"`
	Quantity    int       `db:"quantity"`
	UnitPrice   float64   `db:"unit_price"`
	TotalAmount float64   `db:"total_amount"`
	Status      string    `db:"status"`
	Note        *string   `db:"note"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func OrderToDomain(m Order) orderDom.Order {
	return orderDom.Order{
		ID:          m.ID,
		UserID:      m.UserID,
		ProductID:   m.ProductID,
		Quantity:    m.Quantity,
		UnitPrice:   m.UnitPrice,
		TotalAmount: m.TotalAmount,
		Status:      m.Status,
		Note:        m.Note,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func OrderFromDomain(d orderDom.Order) Order {
	return Order{
		ID:          d.ID,
		UserID:      d.UserID,
		ProductID:   d.ProductID,
		Quantity:    d.Quantity,
		UnitPrice:   d.UnitPrice,
		TotalAmount: d.TotalAmount,
		Status:      d.Status,
		Note:        d.Note,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func ProductToDomain(m Product) orderDom.Product {
	return orderDom.Product{
		ID:          m.ID,
		Category:    m.Category,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Stock:       m.Stock,
		IsActive:    m.IsActive,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ProductFromDomain(d orderDom.Product) Product {
	return Product{
		ID:          d.ID,
		Category:    d.Category,
		Name:        d.Name,
		Description: d.Description,
		Price:       d.Price,
		Stock:       d.Stock,
		IsActive:    d.IsActive,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
