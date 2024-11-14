package composites

import (
	"context"
	"go-clean/pkg/redis"
)

func NewRedis(ctx context.Context, options *redis.Options) (*redis.Redis, error) {
	return redis.New(ctx, options)
}
