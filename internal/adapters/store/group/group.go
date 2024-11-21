package group

import "go-clean/pkg/redis"

type Group struct {
	redis *redis.Redis
}

func New(redis *redis.Redis) *Group {
	return &Group{
		redis: redis,
	}
}
