package service

import (
	"easy-life-back-go/internal/server/routes/auth/redis"
	"easy-life-back-go/internal/server/utils/response"
	"easy-life-back-go/internal/utils"
	"errors"
	"net/http"
)

type Service struct {
	redis *redis.Store
}

func NewService(redis *redis.Store) *Service {
	return &Service{
		redis: redis,
	}
}

func (s *Service) Registration(email string) error {
	randomCode, err := utils.GenerateRandomCode(6)

	if err != nil {
		return err
	}

	err = s.redis.SetRegistrationCode(email, randomCode, 0)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RegistrationSuccess(name, email, password, code string) error {
	val, attempt, err := s.redis.GetRegistrationCode(email)

	if err != nil {
		return errors.New("get code - " + err.Error())
	}

	if attempt > 3 {
		return response.NewBad(http.StatusBadRequest, response.CodeMaxAttemptsExceeded)
	}

	if code != val {
		return response.NewBad(http.StatusBadRequest, response.CodeInvalidCode)
	}

	err = s.redis.DelRegistrationCode(email)

	if err != nil {
		return errors.New("del code - " + err.Error())
	}

	return nil
}
