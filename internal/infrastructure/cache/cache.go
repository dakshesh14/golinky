package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Close() error
}

type redisService struct {
	client *redis.Client
}

func NewCacheService(redisURL string) (CacheService, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &redisService{
		client: client,
	}, nil
}

func (r *redisService) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var strValue string

	switch v := value.(type) {
	case string:
		strValue = v
	default:
		bytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		strValue = string(bytes)
	}

	return r.client.Set(ctx, key, strValue, expiration).Err()
}

func (r *redisService) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisService) Close() error {
	return r.client.Close()
}
