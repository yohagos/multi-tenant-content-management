package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yohagos/multi-content-management/internal/core/port"
)

type cacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) port.CacheRepository {
	return &cacheRepository{
		client: client,
	}
}

func (r *cacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *cacheRepository) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *cacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *cacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r *cacheRepository) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(ctx, key).Result()
}