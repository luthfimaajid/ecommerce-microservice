package cmd

import (
	paymentHndlr "ecom-ms/payment-service/internal/delivery/http"
	"ecom-ms/payment-service/internal/infrastructure/config"
	paymentUc "ecom-ms/payment-service/internal/usecase/payment"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	orderUsecase := paymentUc.NewUsecase()
	orderHandler := paymentHndlr.NewOrderHandler(*orderUsecase)

	r := gin.Default()

	paymentHndlr.RegisterPaymentRoutes(r, orderHandler)

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func newPostgres(cfg config.PostgresConfig) (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
