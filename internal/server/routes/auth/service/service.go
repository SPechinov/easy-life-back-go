package service

import (
	"easy-life-back-go/internal/server/routes/auth/store"
	"easy-life-back-go/internal/server/utils/response"
	"easy-life-back-go/internal/utils"
	"easy-life-back-go/pkg/crypto"
	"errors"
	"fmt"
	"net/http"
)

type Service struct {
	store  *store.Store
	crypto crypto.Crypto
}

func NewService(store *store.Store, crypto crypto.Crypto) *Service {
	return &Service{
		store:  store,
		crypto: crypto,
	}
}

func (s *Service) Registration(email string) error {
	randomCode, err := utils.GenerateRandomCode(6)
	if err != nil {
		return errors.New("generate random code - " + err.Error())
	}

	err = s.store.SetRegistrationCode(email, randomCode, 0)
	if err != nil {
		return errors.New("set code - " + err.Error())
	}

	return nil
}

func (s *Service) RegistrationSuccess(name, email, password, code string) error {
	// Check codes
	has, err := s.store.HasRegistrationCode(email)
	if err != nil {
		return errors.New("has code - " + err.Error())
	}

	if !has {
		return response.NewBad(http.StatusBadRequest, response.CodeDidntSendCode)
	}

	storeCode, gotCount, err := s.store.GetRegistrationCode(email)
	if err != nil {
		return errors.New("get code - " + err.Error())
	}

	if gotCount >= 3 {
		return response.NewBad(http.StatusBadRequest, response.CodeMaxAttemptsExceeded)
	}

	if code != storeCode {
		err = s.store.UpdateGotCountRegistrationCode(email, storeCode, gotCount+1)
		if err != nil {
			return errors.New("update code - " + err.Error())
		}

		return response.NewBad(http.StatusBadRequest, response.CodeInvalidCode)
	}

	err = s.store.DelRegistrationCode(email)
	if err != nil {
		return errors.New("del code - " + err.Error())
	}

	// Encrypt data
	encryptedEmail, err := s.crypto.Encrypt(email)
	if err != nil {
		return errors.New("encrypt email - " + err.Error())
	}

	encryptedPassword, err := s.crypto.Encrypt(email)
	if err != nil {
		return errors.New("encrypt password - " + err.Error())
	}

	fmt.Println(encryptedEmail, encryptedPassword)

	return nil
}
