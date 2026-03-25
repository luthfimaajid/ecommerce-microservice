package order

import "context"

type Cache interface {
	FindProductByID(ctx context.Context, id int64) (*Product, error)
	SetProduct(ctx context.Context, product *Product) error
}
