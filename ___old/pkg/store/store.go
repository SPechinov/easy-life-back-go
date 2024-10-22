package store

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type Store interface {
	Set(key string, value interface{}) error
	SetWithTTL(key string, value any, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	TTL(key string) (time.Duration, error)
	Has(key string) (bool, error)
	Close() error
}

var ctx = context.Background()

type storeClient struct {
	client *redis.Client
}

func New(address, password string, db int) Store {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	client := &storeClient{client: rdb}

	return client
}

func (r *storeClient) Set(key string, value interface{}) error {
	_, err := r.client.Set(ctx, key, value, 0).Result()
	return err
}

func (r *storeClient) SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	_, err := r.client.Set(ctx, key, value, ttl).Result()
	return err
}

func (r *storeClient) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *storeClient) Del(key string) error {
	_, err := r.client.Del(ctx, key).Result()
	return err
}

func (r *storeClient) Has(key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *storeClient) TTL(key string) (time.Duration, error) {
	return r.client.TTL(ctx, key).Result()
}

func (r *storeClient) Close() error {
	err := r.client.Close()

	if err != nil {
		return errors.New("failed to close redis connection: " + err.Error())
	}

	return nil
}
