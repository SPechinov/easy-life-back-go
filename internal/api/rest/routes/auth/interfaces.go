package auth

import (
	"context"
	"go-clean/internal/entities"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type useCases interface {
	Login(ctx context.Context, data entities.UserLogin) (sessionID, accessJWT, refreshJWT string, err error)
	Registration(ctx context.Context, data entities.UserAdd) error
	RegistrationConfirm(ctx context.Context, data entities.UserAddConfirm) error
	ForgotPassword(ctx context.Context, data entities.UserForgotPassword) error
	ForgotPasswordConfirm(ctx context.Context, data entities.UserForgotPasswordConfirm) error
	UpdateJWT(ctx context.Context, userID, sessionID, refreshJWT string) (newSessionID, newAccessJWT, newRefreshJWT string, err error)
	Logout(ctx context.Context, userID, sessionID string)
	LogoutAll(ctx context.Context, userID string)
}
