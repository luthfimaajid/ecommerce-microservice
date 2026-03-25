package order

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindProductByID(ctx context.Context, id int64) (*Product, error)
	InsertOrder(ctx context.Context, tx *sqlx.Tx, d Order) (Order, error)
	CreateTx(ctx context.Context) (*sqlx.Tx, error)
}
