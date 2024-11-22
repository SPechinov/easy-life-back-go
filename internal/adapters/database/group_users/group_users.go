package group_users

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"go-clean/pkg/postgres"
	"time"
)

type GroupUsers struct {
	postgres postgres.Client
}

func New(postgres postgres.Client) *GroupUsers {
	return &GroupUsers{
		postgres: postgres,
	}
}

func (gu *GroupUsers) GetUsersList(ctx context.Context, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	query := `
		SELECT
			public.groups_users.invited_at,
   			public.groups_users.permission,
   			
   			public.users.id,
   			public.users.email,
   			public.users.phone,
   			public.users.first_name,
   			public.users.last_name,
   			public.users.created_at,
   			public.users.updated_at,
   			public.users.deleted_at
		FROM public.groups_users
		
		LEFT JOIN public.users ON public.users.id = public.groups_users.user_id
		
		WHERE public.groups_users.group_id = $1
		
	`

	rows, err := gu.postgres.Query(ctx, query, entity.GroupID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	defer rows.Close()

	users := make([]entities.GroupUser, 0)
	for rows.Next() {
		var user dataUser
		err = rows.Scan(
			&user.invitedAt,
			&user.permission,
			&user.id,
			&user.email,
			&user.phone,
			&user.firstName,
			&user.lastName,
			&user.createdAt,
			&user.updatedAt,
			&user.deletedAt,
		)

		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		users = append(users, entities.GroupUser{
			ID:         user.id,
			Email:      helpers.GetValueFromSQLNullString(user.email),
			Phone:      helpers.GetValueFromSQLNullString(user.phone),
			FirstName:  user.firstName,
			LastName:   helpers.GetPtrValueFromSQLNullString(user.lastName),
			Permission: user.permission,
			CreatedAt:  user.createdAt.Format(time.RFC3339),
			UpdatedAt:  user.updatedAt.Format(time.RFC3339),
			DeletedAt:  helpers.GetPtrValueFromSQLNullTime(user.deletedAt, time.RFC3339),
			InvitedAt:  user.invitedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (gu *GroupUsers) InviteUser(ctx context.Context, entity entities.GroupInviteUser) error {
	query :=
		`
			INSERT INTO public.groups_users (group_id, user_id, permission)
			VALUES ($1, $2, 0)
		`

	_, err := gu.postgres.Exec(ctx, query, entity.GroupID, entity.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return client_error.ErrUserInvited
		}

		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (gu *GroupUsers) ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error {
	query :=
		`
			DELETE FROM public.groups_users WHERE group_id = $1 AND user_id = $2 AND permission != $3
		`

	_, err := gu.postgres.Exec(ctx, query, entity.GroupID, entity.UserID, constants.DefaultAdminPermission)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
