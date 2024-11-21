package codes

import (
	"context"
	"errors"
	"fmt"
	"go-clean/internal/constants"
	"go-clean/pkg/client_error"
	"go-clean/pkg/logger"
	"go-clean/pkg/redis"
	"time"
)

const (
	codeFormat = "%v : %d"
)

type Codes struct {
	redis *redis.Redis
}

func New(redis *redis.Redis) *Codes {
	return &Codes{redis: redis}
}

func (c *Codes) SetCode(ctx context.Context, key, code string, attempts int, ttl time.Duration) error {
	err := c.redis.SetWithTTL(key, fmt.Sprintf(codeFormat, code, attempts), ttl)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (c *Codes) CompareCodes(ctx context.Context, key, code string) error {
	storeCode, attempts, err := c.GetCode(ctx, key)
	if err != nil {
		return err
	}

	if storeCode != code {
		if attempts+1 >= constants.MaxCodeCompareAttempt {
			logger.Debug(ctx, "Max code attempt")
			_ = c.DeleteCode(ctx, key)
			return client_error.ErrCodeMaxAttempts
		}

		logger.Debug(ctx, "Codes not equal")
		_ = c.IncrementCodeAttempts(ctx, key, attempts+1)
		return client_error.ErrCodesIsNotEqual
	}

	_ = c.DeleteCode(ctx, key)

	return nil
}

func (c *Codes) GetCode(ctx context.Context, key string) (code string, attempts int, err error) {
	redisValue, err := c.redis.Get(key)
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return "", 0, client_error.ErrCodeIsNotInRedis
		}
		logger.Error(ctx, err)
		return "", 0, err
	}

	// Parse redis value
	_, err = fmt.Sscanf(redisValue, codeFormat, &code, &attempts)
	if err != nil {
		logger.Error(ctx, err)
		return "", 0, err
	}

	return code, attempts, nil
}

func (c *Codes) IncrementCodeAttempts(ctx context.Context, key string, attempts int) error {
	redisValue, _, err := c.GetCode(ctx, key)
	if err != nil {
		if errors.Is(err, redis.NotFoundError) {
			return client_error.ErrCodeIsNotInRedis
		}
		return fmt.Errorf("update code - " + err.Error())
	}

	ttl, err := c.redis.TTL(key)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = c.SetCode(ctx, key, redisValue, attempts, ttl)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (c *Codes) DeleteCode(ctx context.Context, key string) error {
	err := c.redis.Delete(key)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
