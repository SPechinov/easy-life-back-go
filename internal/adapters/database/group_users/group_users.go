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
		LEFT JOIN public.groups ON public.groups.id = public.groups_users.group_id
		
		WHERE public.groups_users.group_id = $1 AND public.groups.deleted_at IS NULL
		
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

		if user.deletedAt.Valid {
			users = append(
				users,
				entities.GroupUser{
					ID:        user.id,
					DeletedAt: helpers.GetPtrValueFromSQLNullTime(user.deletedAt, time.RFC3339),
				},
			)
		} else {
			users = append(
				users,
				entities.GroupUser{
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
				},
			)
		}

	}

	return users, nil
}

func (gu *GroupUsers) InviteUser(ctx context.Context, entity entities.GroupInviteUser) error {
	query := `
		WITH 
			-- find group
			valid_group AS (
        	    SELECT id
        	    FROM public.groups
        	    WHERE id = $1 AND deleted_at IS NULL
        	),
			-- find user
			valid_user AS (
			    SELECT id
			    FROM public.users
			    WHERE id = $2 AND deleted_at IS NULL
			)
		-- add user in group
		INSERT INTO public.groups_users (group_id, user_id, permission)
		SELECT valid_group.id, valid_user.id, 0
		FROM valid_group CROSS JOIN valid_user
	`

	res, err := gu.postgres.Exec(ctx, query, entity.GroupID, entity.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return client_error.ErrUserNotFound
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return client_error.ErrUserInvited
		}

		logger.Error(ctx, err)
		return err
	}

	if res.RowsAffected() == 0 {
		return client_error.ErrUserNotFound
	}

	return nil
}

func (gu *GroupUsers) ExcludeUser(ctx context.Context, entity entities.GroupExcludeUser) error {
	query := `
		DELETE FROM public.groups_users
		USING public.groups
		WHERE 
			public.groups_users.group_id = $1 
			AND public.groups_users.user_id = $2 
			AND public.groups_users.permission != $3
			AND public.groups.id = public.groups_users.group_id
			AND public.groups.deleted_at IS NULL
	`

	res, err := gu.postgres.Exec(ctx, query, entity.GroupID, entity.UserID, constants.DefaultAdminPermission)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if res.RowsAffected() == 0 {
		return client_error.ErrUserNotFound
	}

	return nil
}
