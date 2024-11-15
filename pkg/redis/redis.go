package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-clean/pkg/helpers"
	"time"
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

	err := connect(ctx, client)
	if err != nil {
		return nil, err
	}

	return &Redis{
		ctx:    ctx,
		client: client,
	}, nil
}

func connect(ctx context.Context, client *redis.Client) error {
	fmt.Println("Redis connecting...")

	err := helpers.Repeatable(
		func() error {
			fmt.Println("Redis try to connect")

			pingErr := client.Ping(ctx).Err()
			if pingErr != nil {
				return pingErr
			}

			return nil
		},
		10,
		2*time.Second,
	)

	if err != nil {
		fmt.Printf("Redis not connected: %s\n", err)
		return err
	}

	fmt.Println("Redis connected")
	return nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}
