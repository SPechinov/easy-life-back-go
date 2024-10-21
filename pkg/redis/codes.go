package redis

import (
	"errors"
	"fmt"
	"time"
)

func (r *redisClient) SetCodeWithTTL(key, code string, attempt int, time time.Duration) error {
	return r.SetWithTTL(key, fmt.Sprintf("%v : %d", code, attempt), time)
}

func (r *redisClient) GetCodeWithTTL(key string) (string, int, error) {
	// Get value
	redisValue, err := r.Get(key)
	if err != nil {
		return "", 0, errors.New("get code - " + err.Error())
	}

	// Get TTL
	ttl, err := r.client.TTL(ctx, key).Result()
	if err != nil {
		return "", 0, err
	}

	// Parse value
	var code string
	var attempt int

	_, err = fmt.Sscanf(redisValue, "%v : %d", &code, &attempt)
	if err != nil {
		return "", 0, errors.New("scan code - " + err.Error())
	}

	attempt += 1

	// Update key record
	err = r.SetCodeWithTTL(key, code, attempt, ttl)
	if err != nil {
		return "", 0, errors.New("set new ttl - " + err.Error())
	}

	return code, attempt, nil
}
