package rest_auth

import (
	"context"
	"go-clean/internal/entities"
	"time"
)

type codes interface {
	SetCode(ctx context.Context, key, code string, attempts int, ttl time.Duration) error
	GetCode(ctx context.Context, key string) (string, int, error)
	CompareCodes(ctx context.Context, key, code string) error
	DeleteCode(ctx context.Context, key string) error
}

type restAuthStore interface {
	SetSession(ctx context.Context, userID, sessionID, refreshJWT string) error
	GetSession(ctx context.Context, userID, sessionID string) (string, error)
	DeleteSession(ctx context.Context, userID, sessionID string)
	DeleteAllSessions(ctx context.Context, userID string)
}

type usersService interface {
	Add(ctx context.Context, data entities.UserAddConfirm) error
	Get(ctx context.Context, data entities.UserGet) (*entities.User, error)
	Restore(ctx context.Context, data entities.UserAddConfirm) error
	UpdatePassword(ctx context.Context, data entities.UserForgotPasswordConfirm) error
	GetDeletedTime(ctx context.Context, entity entities.UserGet) (*time.Time, error)
}
