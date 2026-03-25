package inventory

import (
	"context"
	domain "ecom-ms/inventory-service/internal/domain/inventory"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	db     *sqlx.DB
	schema string
}

func NewPostgresRepository(db *sqlx.DB, schema string) *PostgresRepository {
	return &PostgresRepository{
		db:     db,
		schema: schema,
	}
}

func (r *PostgresRepository) CreateTx(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.BeginTxx(ctx, nil)
}

func (r *PostgresRepository) InsertStockHistory(
	ctx context.Context,
	tx *sqlx.Tx,
	s *domain.StockHistory,
) error {
	stockHistory := FromDomain(*s)

	query := `
	INSERT INTO inventory.stock_history (
		product_id,
		order_id,
		user_id,
		movement_type,
		quantity,
		created_at
	)
	VALUES (
		:product_id,
		:order_id,
		:user_id,
		:movement_type,
		:quantity,
		:created_at
	)
	RETURNING id
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.GetContext(ctx, &stockHistory, stockHistory); err != nil {
		return err
	}

	return nil
}
