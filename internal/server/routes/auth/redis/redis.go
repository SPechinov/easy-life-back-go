package redis

import (
	pkgRedis "easy-life-back-go/pkg/redis"
	"time"
)

type Redis struct {
	redis pkgRedis.Client
}

func NewRedis(redis pkgRedis.Client) *Redis {
	return &Redis{
		redis: redis,
	}
}

func (a *Redis) SetRegistrationCode(email, code string) error {
	return a.redis.SetWithTTL(GetKeyHttpUserRegistrationCode(email), code, time.Minute*10)
}

func (a *Redis) SetForgotRegistrationCode(email, code string) error {
	return nil
}

func (a *Redis) SetJWTPairCode(id, code string) error {
	return nil
}
