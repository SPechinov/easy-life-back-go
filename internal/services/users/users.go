package users

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"time"
)

type database interface {
	Add(ctx context.Context, data entities.UserAddConfirm) error
	Restore(ctx context.Context, data entities.UserAddConfirm) error
	Get(ctx context.Context, data entities.UserGet) (*entities.User, error)
	UpdatePassword(ctx context.Context, data entities.UserForgotPasswordConfirm) error
	GetDeletedTime(ctx context.Context, entity entities.UserGet) (*time.Time, error)
}

type Users struct {
	database database
}

func New(database database) *Users {
	return &Users{
		database: database,
	}
}

func (u *Users) Add(ctx context.Context, data entities.UserAddConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = u.database.Add(ctx, entities.UserAddConfirm{
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

	err = u.database.Restore(ctx, entities.UserAddConfirm{
		AuthWay:   data.AuthWay,
		FirstName: data.FirstName,
		Password:  hashedPassword,
	})

	return err
}

func (u *Users) Get(ctx context.Context, data entities.UserGet) (*entities.User, error) {
	user, err := u.database.Get(ctx, data)
	return user, err
}

func (u *Users) GetDeletedTime(ctx context.Context, entity entities.UserGet) (*time.Time, error) {
	return u.database.GetDeletedTime(ctx, entity)
}

func (u *Users) UpdatePassword(ctx context.Context, data entities.UserForgotPasswordConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	err = u.database.UpdatePassword(ctx, entities.UserForgotPasswordConfirm{
		AuthWay:  data.AuthWay,
		Password: hashedPassword,
		Code:     data.Code,
	})
	return err
}
