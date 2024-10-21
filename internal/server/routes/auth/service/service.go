package service

import (
	"easy-life-back-go/internal/server/routes/auth/store"
	"easy-life-back-go/internal/server/utils/response"
	"easy-life-back-go/internal/utils"
	"errors"
	"net/http"
)

type Service struct {
	store *store.Store
}

func NewService(redis *store.Store) *Service {
	return &Service{
		store: redis,
	}
}

func (s *Service) Registration(email string) error {
	randomCode, err := utils.GenerateRandomCode(6)

	if err != nil {
		return err
	}

	err = s.store.SetRegistrationCode(email, randomCode, 0)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RegistrationSuccess(name, email, password, code string) error {
	val, attempt, err := s.store.GetRegistrationCode(email)

	if err != nil {
		return errors.New("get code - " + err.Error())
	}

	if attempt > 3 {
		return response.NewBad(http.StatusBadRequest, response.CodeMaxAttemptsExceeded)
	}

	if code != val {
		return response.NewBad(http.StatusBadRequest, response.CodeInvalidCode)
	}

	err = s.store.DelRegistrationCode(email)

	if err != nil {
		return errors.New("del code - " + err.Error())
	}

	return nil
}
