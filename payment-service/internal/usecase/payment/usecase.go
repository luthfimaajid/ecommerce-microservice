package payment

import (
	"context"
	"errors"
	"math/rand"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) CreatePayment(
	ctx context.Context,
	userId int64,
	totalAmount float64,
) error {

	if rand.Intn(2) == 0 {
		return errors.New("payment failed")
	}

	return nil
}
