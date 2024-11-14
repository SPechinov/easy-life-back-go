package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
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
	connectionErr := client.Ping(ctx).Err()

	if connectionErr == nil {
		fmt.Println("Redis connected")
		return nil
	}

	tryCount := 1
	for tryCount < 10 {
		time.Sleep(2 * time.Second)

		fmt.Printf("Redis try to connect: %d time\n", tryCount+1)
		tryCount++
		connectionErr = client.Ping(ctx).Err()

		if connectionErr == nil {
			fmt.Println("Redis connected")
			return nil
		}
	}

	fmt.Printf("Redis not connected: %s\n", connectionErr)
	return connectionErr
}

func (r *Redis) Close() error {
	return r.client.Close()
}
