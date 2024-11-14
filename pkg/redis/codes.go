package redis

import (
	"errors"
	"fmt"
	"time"
)

const (
	codeFormat = "%v : %d"
)

// SetCode use for storing codes like email, sms and so on
func (r *Redis) SetCode(key, code string, attempts int, ttl time.Duration) error {
	err := r.SetWithTTL(key, fmt.Sprintf(codeFormat, code, attempts), ttl)
	if err != nil {
		return fmt.Errorf("set code - " + err.Error())
	}

	return nil
}

// GetCode can return NotFoundError
func (r *Redis) GetCode(key string) (code string, attempts int, err error) {
	redisValue, err := r.Get(key)
	if err != nil {
		if errors.Is(err, NotFoundError) {
			return "", 0, NotFoundError
		}
		return "", 0, fmt.Errorf("get code - " + err.Error())
	}

	// Parse redis value
	_, err = fmt.Sscanf(redisValue, codeFormat, &code, &attempts)
	if err != nil {
		return "", 0, fmt.Errorf("get code - " + err.Error())
	}

	return code, attempts, nil
}

// UpdateCode can return NotFoundError
func (r *Redis) UpdateCode(key string, attempts int) error {
	redisValue, _, err := r.GetCode(key)
	if err != nil {
		if errors.Is(err, NotFoundError) {
			return NotFoundError
		}
		return fmt.Errorf("update code - " + err.Error())
	}

	ttl, err := r.TTL(key)
	if err != nil {
		return fmt.Errorf("update code - " + err.Error())
	}

	err = r.SetCode(key, redisValue, attempts, ttl)
	if err != nil {
		return fmt.Errorf("update code - " + err.Error())
	}

	return nil
}
