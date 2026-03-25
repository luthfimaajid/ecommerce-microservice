package order

import (
	"context"
	domain "ecom-ms/order-service/internal/domain/order"
	orderDom "ecom-ms/order-service/internal/domain/order"
	"fmt"

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

func (r *PostgresRepository) FindProductByID(ctx context.Context, id int64) (*domain.Product, error) {
	var m Product

	query := fmt.Sprintf(`
	SELECT 
		id, 
		category, 
		name, 
		description, 
		price, 
		stock, 
		is_active, 
		created_at, 
		updated_at
	FROM "%[1]s".product
	WHERE 
		id = $1
	`, r.schema)

	err := r.db.GetContext(ctx, &m, query, id)
	if err != nil {
		return nil, err
	}

	d := ProductToDomain(m)

	return &d, nil
}

func (r *PostgresRepository) InsertOrder(
	ctx context.Context,
	tx *sqlx.Tx,
	d orderDom.Order,
) (orderDom.Order, error) {
	model := OrderFromDomain(d)

	query := fmt.Sprintf(`
	INSERT INTO "%[1]s"."order" (
		user_id,
		product_id,
		quantity,
		unit_price,
		total_amount,
		status,
		note
	)
	VALUES (
		:user_id,
		:product_id,
		:quantity,
		:unit_price,
		:total_amount,
		:status,
		:note
	)
	RETURNING id, created_at, updated_at
	`, r.schema)

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return orderDom.Order{}, err
	}
	defer stmt.Close()

	if err := stmt.GetContext(ctx, &model, model); err != nil {
		return orderDom.Order{}, err
	}

	return OrderToDomain(model), nil
}
