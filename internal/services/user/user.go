package user

import (
	"context"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
)

type database interface {
	AddUser(ctx context.Context, data entities.UserAddConfirm) error
	RestoreUser(ctx context.Context, data entities.UserAddConfirm) error
	GetUser(ctx context.Context, data entities.UserGet) (*entities.User, error)
	UpdatePasswordUser(ctx context.Context, data entities.UserForgotPasswordConfirm) error
}

type User struct {
	database database
}

func New(database database) *User {
	return &User{
		database: database,
	}
}

func (u *User) AddUser(ctx context.Context, data entities.UserAddConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = u.database.AddUser(ctx, entities.UserAddConfirm{
		AuthWay:   data.AuthWay,
		FirstName: data.FirstName,
		Password:  hashedPassword,
		Code:      data.Code,
	})

	return err
}

func (u *User) RestoreUser(ctx context.Context, data entities.UserAddConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	err = u.database.RestoreUser(ctx, entities.UserAddConfirm{
		AuthWay:   data.AuthWay,
		FirstName: data.FirstName,
		Password:  hashedPassword,
	})

	return err
}

func (u *User) GetUser(ctx context.Context, data entities.UserGet) (*entities.User, error) {
	user, err := u.database.GetUser(ctx, data)
	return user, err
}

func (u *User) UpdatePasswordUser(ctx context.Context, data entities.UserForgotPasswordConfirm) error {
	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}
	err = u.database.UpdatePasswordUser(ctx, entities.UserForgotPasswordConfirm{
		AuthWay:  data.AuthWay,
		Password: hashedPassword,
		Code:     data.Code,
	})
	return err
}
