package cmd

import (
	orderHndlr "ecom-ms/order-service/internal/delivery/http"
	"ecom-ms/order-service/internal/infrastructure/config"
	orderInfra "ecom-ms/order-service/internal/infrastructure/order"
	orderUc "ecom-ms/order-service/internal/usecase/order"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr(),
		Password: cfg.Redis.Password,
	})

	nats, err := nats.Connect(cfg.Nats.URL())

	orderRepo := orderInfra.NewPostgresRepository(db, cfg.Postgres.Schema)
	orderMessage := orderInfra.NewNatsPublisher(nats)
	orderCache := orderInfra.NewRedisRepository(redisClient)
	orderUsecase := orderUc.NewUsecase(orderRepo, orderCache, orderMessage)

	orderHandler := orderHndlr.NewOrderHandler(*orderUsecase)

	r := gin.Default()

	orderHndlr.RegisterOrderRoutes(r, orderHandler)

	if err := r.Run(":8080"); err != nil {
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
