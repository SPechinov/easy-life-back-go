package service

import (
	"easy-life-back-go/internal/server/routes/auth/redis"
	"easy-life-back-go/internal/server/utils/response"
	"errors"
)

type Service struct {
	redis *redis.Redis
}

func NewService(redis *redis.Redis) *Service {
	return &Service{
		redis: redis,
	}
}

func (s *Service) Registration(name, email, password string) error {
	err := s.redis.SetRegistrationCode(email, "4444")

	if err != nil {
		return errors.New(response.SomethingHappen)
	}

	return nil
}
