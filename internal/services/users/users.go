package users

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"time"
)

type usersDatabase interface {
	Add(ctx context.Context, data entities.UserAddConfirm) error
	Restore(ctx context.Context, data entities.UserAddConfirm) error
	Get(ctx context.Context, data entities.UserGet) (*entities.User, error)
	UpdatePassword(ctx context.Context, data entities.UserForgotPasswordConfirm) error
	GetDeletedTime(ctx context.Context, entity entities.UserGet) (*time.Time, error)
}

type Users struct {
	usersDatabase usersDatabase
}

func New(database usersDatabase) *Users {
	return &Users{
		usersDatabase: database,
	}
}

func (u *Users) Add(ctx context.Context, data entities.UserAddConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = u.usersDatabase.Add(ctx, entities.UserAddConfirm{
		AuthWay:   data.AuthWay,
		FirstName: data.FirstName,
		Password:  hashedPassword,
		Code:      data.Code,
	})

	return err
}

func (u *Users) Restore(ctx context.Context, data entities.UserAddConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = u.usersDatabase.Restore(ctx, entities.UserAddConfirm{
		AuthWay:   data.AuthWay,
		FirstName: data.FirstName,
		Password:  hashedPassword,
	})

	return err
}

func (u *Users) Get(ctx context.Context, data entities.UserGet) (*entities.User, error) {
	user, err := u.usersDatabase.Get(ctx, data)
	return user, err
}

func (u *Users) GetDeletedTime(ctx context.Context, entity entities.UserGet) (*time.Time, error) {
	return u.usersDatabase.GetDeletedTime(ctx, entity)
}

func (u *Users) UpdatePassword(ctx context.Context, data entities.UserForgotPasswordConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	err = u.usersDatabase.UpdatePassword(ctx, entities.UserForgotPasswordConfirm{
		AuthWay:  data.AuthWay,
		Password: hashedPassword,
		Code:     data.Code,
	})
	return err
}
