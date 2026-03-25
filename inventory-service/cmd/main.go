package cmd

import (
	"ecom-ms/inventory-service/internal/infrastructure/config"
	inventoryInfra "ecom-ms/inventory-service/internal/infrastructure/inventory"
	inventoryUc "ecom-ms/inventory-service/internal/usecase/inventory"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := newPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nats, err := nats.Connect(cfg.Nats.URL())

	inventoryRepo := inventoryInfra.NewPostgresRepository(db, cfg.Postgres.Schema)
	inventoryUsecase := inventoryUc.NewOrderEventHandler(inventoryRepo)
	inventoryConsumer := inventoryInfra.NewNatsConsumer(nats, inventoryUsecase)

	err = inventoryConsumer.Start()
	if err != nil {
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
