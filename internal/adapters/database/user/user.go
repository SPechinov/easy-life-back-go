package user

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
	"go-clean/pkg/logger"
	"go-clean/pkg/postgres"
	"time"
)

type User struct {
	postgres *postgres.Postgres
}

func New(postgres *postgres.Postgres) *User {
	return &User{
		postgres: postgres,
	}
}

type authData struct {
	email *string
	phone *string
}

func getAuthData(authWay entities.UserAuthWay) authData {
	var email *string
	if authWay.Email != "" {
		email = &authWay.Email
	}

	var phone *string
	if authWay.Phone != "" {
		phone = &authWay.Phone
	}

	return authData{
		email: email,
		phone: phone,
	}
}

func (u *User) AddUser(ctx context.Context, data entities.UserAddConfirm) error {
	ad := getAuthData(data.AuthWay)

	query := `
		INSERT INTO users (email, phone, first_name, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := u.postgres.Exec(query, ad.email, ad.phone, data.FirstName, data.Password)

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (u *User) RestoreUser(ctx context.Context, data entities.UserAddConfirm) error {
	ad := getAuthData(data.AuthWay)

	query := `UPDATE users SET first_name = $1, password = $2, deleted_at = null WHERE email = $3 OR phone = $4`

	_, err := u.postgres.Exec(query, data.FirstName, data.Password, ad.email, ad.phone)

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (u *User) GetUser(ctx context.Context, data entities.UserGet) (*entities.User, error) {
	if data.ID == "" && data.Email == "" && data.Phone == "" {
		return nil, client_error.ErrUserNotFound
	}

	// Create query
	var userData dataUser
	var err error
	if data.ID == "" {
		query := `SELECT id, email, phone, password, first_name, last_name, created_at, updated_at, deleted_at FROM users WHERE email = $1 OR phone = $2`
		err = u.postgres.QueryRow(query, data.Email, data.Phone).Scan(
			&userData.ID,
			&userData.Email,
			&userData.Phone,
			&userData.Password,
			&userData.FirstName,
			&userData.LastName,
			&userData.CreatedAt,
			&userData.UpdatedAt,
			&userData.DeletedAt,
		)
	} else {
		query := `SELECT id, email, phone, password, first_name, last_name, created_at, updated_at, deleted_at FROM users WHERE id = $1`
		err = u.postgres.QueryRow(query, data.ID).Scan(
			&userData.ID,
			&userData.Email,
			&userData.Phone,
			&userData.Password,
			&userData.FirstName,
			&userData.LastName,
			&userData.CreatedAt,
			&userData.UpdatedAt,
			&userData.DeletedAt,
		)
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, client_error.ErrUserNotFound
		}
		logger.Error(ctx, err)
		return nil, err
	}

	// Parse data
	user := new(entities.User)
	user.ID = userData.ID
	if userData.Email.Valid {
		user.Email = userData.Email.String
	}
	if userData.Phone.Valid {
		user.Phone = userData.Phone.String
	}
	user.Password = string(userData.Password)
	user.FirstName = userData.FirstName
	if userData.LastName.Valid {
		user.LastName = &userData.LastName.String
	}
	user.CreatedAt = userData.CreatedAt.Format(time.RFC3339)
	user.UpdatedAt = userData.UpdatedAt.Format(time.RFC3339)
	if userData.DeletedAt.Valid {
		deletedAtStr := userData.DeletedAt.Time.Format(time.RFC3339)
		user.DeletedAt = &deletedAtStr
	}

	return user, nil
}

func (u *User) UpdatePasswordUser(ctx context.Context, data entities.UserForgotPasswordConfirm) error {
	ad := getAuthData(data.AuthWay)

	query := `UPDATE users SET password = $1 WHERE email = $2 OR phone = $3`

	_, err := u.postgres.Exec(query, data.Password, ad.email, ad.phone)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return client_error.ErrUserNotFound
		}
		logger.Error(ctx, err)
		return err
	}

	return nil
}