package rest_auth

import (
	"context"
	"errors"
	"go-clean/pkg/client_error"
	"go-clean/pkg/logger"
	"go-clean/pkg/redis"
	"time"
)

func getKeyUserForgotPasswordCode(key string) string {
	return "http:rest-auth:forgot-password-code:" + key
}

func (ra *RestAuth) SetForgotPasswordCode(ctx context.Context, key, code string) error {
	err := ra.redis.SetCode(getKeyUserForgotPasswordCode(key), code, 0, 10*time.Minute)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	return nil
}

func (ra *RestAuth) GetForgotPasswordCode(ctx context.Context, key string) (string, int, error) {
	code, attempts, err := ra.redis.GetCode(getKeyUserForgotPasswordCode(key))
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return "", 0, client_error.ErrCodeIsNotInRedis
		}

		logger.Error(ctx, err)
		return "", 0, err
	}

	return code, attempts, nil
}

func (ra *RestAuth) UpdateForgotPasswordCode(ctx context.Context, key string, attempts int) error {
	err := ra.redis.UpdateCode(getKeyUserForgotPasswordCode(key), attempts)
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return client_error.ErrCodeIsNotInRedis
		}

		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (ra *RestAuth) DeleteForgotPasswordCode(ctx context.Context, key string) error {
	err := ra.redis.Delete(getKeyUserForgotPasswordCode(key))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
