package rest_auth

import (
	"context"
	"go-clean/internal/entities"
)

type store interface {
	SetRegistrationCode(ctx context.Context, key, code string) error
	GetRegistrationCode(ctx context.Context, key string) (string, int, error)
	UpdateRegistrationCode(ctx context.Context, key string, attempts int) error
	DeleteRegistrationCode(ctx context.Context, key string) error

	SetForgotPasswordCode(ctx context.Context, key, code string) error
	GetForgotPasswordCode(ctx context.Context, key string) (string, int, error)
	UpdateForgotPasswordCode(ctx context.Context, key string, attempts int) error
	DeleteForgotPasswordCode(ctx context.Context, key string) error

	SetSession(ctx context.Context, userID, sessionID, refreshJWT string) error
	GetSession(ctx context.Context, userID, sessionID string) (string, error)
	DeleteSession(ctx context.Context, userID, sessionID string)
	DeleteAllSessions(ctx context.Context, userID string)
}

type service interface {
	AddUser(ctx context.Context, data entities.UserAddConfirm) error
	GetUser(ctx context.Context, data entities.UserGet) (*entities.User, error)
	RestoreUser(ctx context.Context, data entities.UserAddConfirm) error
	UpdatePasswordUser(ctx context.Context, data entities.UserForgotPasswordConfirm) error
}
