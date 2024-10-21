package store_codes

import (
	"easy-life-back-go/pkg/store"
	"errors"
	"fmt"
	"time"
)

const codeFormat = "%v : %d"

type StoreCodes interface {
	SetWithTTL(key, code string, gotCount int, time time.Duration) error
	GetWithTTL(key string) (string, int, error)
	UpdateGotCountWithTTL(key, code string, gotCount int) error
}

type storeCodes struct {
	client store.Store
}

func NewStoreCodes(client store.Store) StoreCodes {
	return &storeCodes{client: client}
}

func (c *storeCodes) SetWithTTL(key, code string, gotCount int, time time.Duration) error {
	err := c.client.SetWithTTL(key, fmt.Sprintf(codeFormat, code, gotCount), time)

	if err != nil {
		return errors.New("set code ttl - " + err.Error())
	}

	return nil
}

func (c *storeCodes) GetWithTTL(key string) (string, int, error) {
	// Get value
	storeValue, err := c.client.Get(key)
	if err != nil {
		return "", 0, errors.New("get code ttl - " + err.Error())
	}

	// Parse value
	var code string
	var gotCount int

	_, err = fmt.Sscanf(storeValue, codeFormat, &code, &gotCount)
	if err != nil {
		return "", 0, errors.New("scan code ttl - " + err.Error())
	}

	return code, gotCount, nil
}

func (c *storeCodes) UpdateGotCountWithTTL(key, code string, gotCount int) error {
	// Get TTL
	ttl, err := c.client.TTL(key)
	if err != nil {
		return errors.New("get ttl - " + err.Error())
	}

	// Update key record
	err = c.SetWithTTL(key, code, gotCount, ttl)
	if err != nil {
		return errors.New("update gotCount - " + err.Error())
	}

	return nil
}
