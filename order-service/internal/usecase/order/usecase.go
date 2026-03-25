package order

import (
	"context"
	"ecom-ms/order-service/internal/domain/order"
	"time"
)

type Usecase struct {
	repo    order.Repository
	cache   order.Cache
	message order.Messaging
}

func NewUsecase(
	repo order.Repository,
	cache order.Cache,
	message order.Messaging,
) *Usecase {
	return &Usecase{
		repo:    repo,
		cache:   cache,
		message: message,
	}
}

func (u *Usecase) CreateOrder(
	ctx context.Context,
	productID int64,
	quantity int,
	note *string,
) error {

	userID := int64(1)

	tx, err := u.repo.CreateTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	product, err := u.cache.FindProductByID(ctx, productID)
	if err != nil {
		return err
	}

	if product == nil {
		product, err = u.repo.FindProductByID(ctx, productID)
		if err != nil {
			return err
		}

		go u.cache.SetProduct(ctx, product)
	}

	now := time.Now()

	o := order.Order{
		UserID:      userID,
		ProductID:   productID,
		Quantity:    quantity,
		UnitPrice:   product.Price,
		TotalAmount: product.Price * float64(quantity),
		Status:      "pending",
		Note:        note,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	o, err = u.repo.InsertOrder(ctx, tx, o)
	if err != nil {
		return err
	}

	event := order.OrderCreatedEvent{
		UserID:    o.UserID,
		OrderID:   o.ID,
		ProductID: o.ProductID,
		Quantity:  o.Quantity,
	}

	if err := u.message.PublishOrderCreated(ctx, event); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
