package redis

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Get key from store. If it has not got any key return NotFoundError
func (r *Redis) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return "", NotFoundError
	}

	if err != nil {
		return "", fmt.Errorf("could not get key %s: %v", key, err)
	}

	return val, nil
}

// Set key in store
func (r *Redis) Set(key string, value any) error {
	if err := r.client.Set(r.ctx, key, value, 0).Err(); err != nil {
		return fmt.Errorf("could not set key %s: %v", key, err)
	}
	return nil
}

// SetWithTTL set key in store with TTL
func (r *Redis) SetWithTTL(key string, value any, ttl time.Duration) error {
	if err := r.client.Set(r.ctx, key, value, ttl).Err(); err != nil {
		return fmt.Errorf("could not set with ttl key %s: %v", key, err)
	}
	return nil
}

// Delete key from store
func (r *Redis) Delete(keys ...string) error {
	if err := r.client.Del(r.ctx, keys...).Err(); err != nil {
		return fmt.Errorf("could not delete keys %s: %v", keys, err)
	}
	return nil
}

// Exist check has key in store or not
func (r *Redis) Exist(key string) (bool, error) {
	count, err := r.client.Exists(r.ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("could not check existence of key %s: %w", key, err)
	}
	return count > 0, nil
}

// TTL get
func (r *Redis) TTL(key string) (time.Duration, error) {
	ttl, err := r.client.TTL(r.ctx, key).Result()
	if err != nil {
		return time.Duration(0), fmt.Errorf("could not check TTL of key %s: %w", key, err)
	}
	return ttl, nil
}

func (r *Redis) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	keys, newCursor, err := r.client.Scan(r.ctx, cursor, match, count).Result()
	if err != nil {
		return nil, 0, fmt.Errorf("could not scan %s: %w", match, err)
	}
	return keys, newCursor, err
}
