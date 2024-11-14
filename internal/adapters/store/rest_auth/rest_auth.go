package rest_auth

import (
	"go-clean/pkg/redis"
)

type RestAuth struct {
	redis *redis.Redis
}

func New(redis *redis.Redis) RestAuth {
	return RestAuth{
		redis: redis,
	}
}
