package store_codes

import (
	"easy-life-back-go/pkg/store"
	"errors"
	"fmt"
	"time"
)

type StoreCodes interface {
	SetWithTTL(key, code string, attempt int, time time.Duration) error
	GetWithTTL(key string) (string, int, error)
}

type storeCodes struct {
	client store.Store
}

func NewStoreCodes(client store.Store) StoreCodes {
	return &storeCodes{client: client}
}

func (c *storeCodes) SetWithTTL(key, code string, attempt int, time time.Duration) error {
	return c.client.SetWithTTL(key, fmt.Sprintf("%v : %d", code, attempt), time)
}

// TODO: c.client.Get(key) и c.client.TTL(key) можно вызывать одновременно, но я пока не знаю как :)

func (c *storeCodes) GetWithTTL(key string) (string, int, error) {
	// Get value
	redisValue, err := c.client.Get(key)
	if err != nil {
		return "", 0, errors.New("get code - " + err.Error())
	}

	// Get TTL
	ttl, err := c.client.TTL(key)
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
	err = c.SetWithTTL(key, code, attempt, ttl)
	if err != nil {
		return "", 0, errors.New("set new ttl - " + err.Error())
	}

	return code, attempt, nil
}
