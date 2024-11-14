package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

type Options struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func New(ctx context.Context, options *Options) (*Redis, error) {
	var client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", options.Host, options.Port),
		Password: options.Password,
		DB:       options.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("could not connect to redis: %v", err)
	}

	return &Redis{
		ctx:    ctx,
		client: client,
	}, nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}
