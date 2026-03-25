package inventory

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateTx(ctx context.Context) (*sqlx.Tx, error)
	InsertStockHistory(ctx context.Context, tx *sqlx.Tx, s *StockHistory) error
}
