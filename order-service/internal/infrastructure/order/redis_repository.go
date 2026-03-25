package order

import (
	"context"
	domain "ecom-ms/order-service/internal/domain/order"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
	prefix string
	ttl    time.Duration
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
		prefix: "order:product",
		ttl:    10 * time.Minute,
	}
}

func (r *RedisRepository) FindProductByID(ctx context.Context, id int64) (*domain.Product, error) {
	key := fmt.Sprintf("%s:%d", r.prefix, id)

	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var product domain.Product
	if err := json.Unmarshal([]byte(val), &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *RedisRepository) SetProduct(ctx context.Context, product *domain.Product) error {
	key := fmt.Sprintf("%s:%d", r.prefix, product.ID)

	b, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, b, r.ttl).Err()
}
