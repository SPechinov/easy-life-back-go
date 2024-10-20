package service

import "easy-life-back-go/internal/server/routes/auth/redis"

type Service struct {
	redis redis.AuthRedis
}

func NewService(redis redis.AuthRedis) *Service {
	return &Service{
		redis: redis,
	}
}
