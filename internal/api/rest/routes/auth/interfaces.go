package auth

import (
	"context"
	"go-clean/internal/entities"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type useCases interface {
	Login(ctx context.Context, entity entities.UserLogin) (sessionID, accessJWT, refreshJWT string, err error)
	Registration(ctx context.Context, entity entities.UserAdd) error
	RegistrationConfirm(ctx context.Context, entity entities.UserAddConfirm) error
	ForgotPassword(ctx context.Context, entity entities.UserForgotPassword) error
	ForgotPasswordConfirm(ctx context.Context, entity entities.UserForgotPasswordConfirm) error
	UpdateJWT(ctx context.Context, entity entities.UserUpdateJWT) (newSessionID, newAccessJWT, newRefreshJWT string, err error)
	Logout(ctx context.Context, entity entities.UserLogout)
	LogoutAll(ctx context.Context, entity entities.UserLogoutAll)
}
