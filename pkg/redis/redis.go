package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client interface {
	Set(key string, value interface{}) error
	SetWithTTL(key string, value interface{}, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	SetCodeWithTTL(key, code string, attempt int, time time.Duration) error
	GetCodeWithTTL(key string) (string, int, error)
}

var ctx = context.Background()

type redisClient struct {
	client *redis.Client
}

func NewClient(address, password string, db int) Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return &redisClient{client: rdb}
}

func (r *redisClient) Set(key string, value interface{}) error {
	_, err := r.client.Set(ctx, key, value, 0).Result()
	return err
}

func (r *redisClient) SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	_, err := r.client.Set(ctx, key, value, ttl).Result()
	return err
}

func (r *redisClient) Get(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	return val, err
}

func (r *redisClient) Del(key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}
