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

func (r *Redis) SetRegistrationCode(email, code string, attempt int) error {
	return r.redis.SetCodeWithTTL(
		GetKeyUserRegistrationCode(email),
		code,
		attempt,
		time.Minute*10,
	)
}

func (r *Redis) GetRegistrationCode(email string) (string, int, error) {
	return r.redis.GetCodeWithTTL(GetKeyUserRegistrationCode(email))
}

func (r *Redis) DelRegistrationCode(email string) error {
	return r.redis.Del(GetKeyUserRegistrationCode(email))
}

func (r *Redis) SetForgotRegistrationCode(email, code string) error {
	return nil
}

func (r *Redis) SetJWTPairCode(id, code string) error {
	return nil
}
