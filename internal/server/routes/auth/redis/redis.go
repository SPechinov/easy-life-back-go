package redis

import "easy-life-back-go/pkg/redis"

type AuthRedis interface {
	SetRegistrationCode(code string) error
	SetForgotRegistrationCode(code string) error
	SetJWTPairCode(code string) error
}

type authRedis struct {
	redis *redis.Client
}

func NewRedis(redis *redis.Client) AuthRedis {
	return &authRedis{
		redis: redis,
	}
}

func (a *authRedis) SetRegistrationCode(code string) error {
	return nil
}

func (a *authRedis) SetForgotRegistrationCode(code string) error {
	return nil
}

func (a *authRedis) SetJWTPairCode(code string) error {
	return nil
}
