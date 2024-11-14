package rest_auth

import (
	"context"
	"errors"
	"go-clean/pkg/client_error"
	"go-clean/pkg/logger"
	"go-clean/pkg/redis"
	"time"
)

func getKeyUserRegistrationCode(key string) string {
	return "http:rest-auth:reg-code:" + key
}

func (ra *RestAuth) SetRegistrationCode(ctx context.Context, key, code string) error {
	err := ra.redis.SetCode(getKeyUserRegistrationCode(key), code, 0, 10*time.Minute)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

func (ra *RestAuth) GetRegistrationCode(ctx context.Context, key string) (string, int, error) {
	code, attempts, err := ra.redis.GetCode(getKeyUserRegistrationCode(key))
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return "", 0, client_error.ErrCodeIsNotInRedis
		}

		logger.Error(ctx, err)
		return "", 0, err
	}

	return code, attempts, nil
}

func (ra *RestAuth) UpdateRegistrationCode(ctx context.Context, key string, attempts int) error {
	err := ra.redis.UpdateCode(getKeyUserRegistrationCode(key), attempts)
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return client_error.ErrCodeIsNotInRedis
		}

		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (ra *RestAuth) DeleteRegistrationCode(ctx context.Context, key string) error {
	err := ra.redis.Delete(getKeyUserRegistrationCode(key))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
