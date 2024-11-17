package group

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"go-clean/pkg/postgres"
	"time"
)

type Group struct {
	postgres postgres.Client
}

func New(postgres postgres.Client) *Group {
	return &Group{
		postgres: postgres,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (string, error) {
	tx, err := g.postgres.Begin(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	queryAddGroup := `
		INSERT INTO public.groups (name)
		VALUES ($1)
		RETURNING id
	`

	var groupID string
	err = tx.QueryRow(ctx, queryAddGroup, entity.Name).Scan(&groupID)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	queryUsersGroup := `
		INSERT INTO public.users_groups (group_id, user_id, permission)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(ctx, queryUsersGroup, groupID, entity.AdminID, constants.DefaultAdminPermission)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return groupID, nil
}

func (g *Group) Patch(ctx context.Context, entity entities.GroupPatch) error {
	query := `
		UPDATE public.groups
		SET name = COALESCE(NULLIF($1, ''), name)
		WHERE id = $2
	`

	_, err := g.postgres.Exec(ctx, query, entity.Name, entity.GroupID)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
func (g *Group) Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error) {
	query := `
		SELECT
		    -- Group
			public.groups.id AS group_id,
			public.groups.name AS group_name,
			public.groups.is_payed AS group_is_payed,
			public.groups.created_at AS group_created_at,
			public.groups.updated_at AS group_updated_at,
			public.groups.deleted_at AS group_deleted_at,
			-- Admin
			public.users_groups.invited_at AS admin_invited_ad,
			public.users.id AS admin_id,
			public.users.email AS admin_email,
			public.users.phone AS admin_phone,
			public.users.password AS admin_password,
			public.users.first_name AS admin_first_name,
			public.users.last_name AS admin_last_name,
			public.users.created_at AS admin_created_at,
			public.users.updated_at AS admin_updated_at,
			public.users.deleted_at AS admin_deleted_at
		FROM public.groups
		
		-- Admin
		LEFT JOIN public.users_groups
			   ON public.users_groups.group_id = public.groups.id AND public.users_groups.permission = $2
		LEFT JOIN public.users
			   ON public.users.id = public.users_groups.user_id
		
		WHERE public.groups.id = $1
	`

	group := new(dataGroupWithAdmin)
	err := g.postgres.QueryRow(ctx, query, entity.GroupID, constants.DefaultAdminPermission).Scan(
		&group.id,
		&group.name,
		&group.isPayed,
		&group.createdAt,
		&group.updatedAt,
		&group.deletedAt,
		&group.admin.invitedAt,
		&group.admin.id,
		&group.admin.email,
		&group.admin.phone,
		&group.admin.password,
		&group.admin.firstName,
		&group.admin.lastName,
		&group.admin.createdAt,
		&group.admin.updatedAt,
		&group.admin.deletedAt,
	)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &entities.Group{
		ID:        group.id,
		Name:      group.name,
		IsPayed:   group.isPayed,
		CreatedAt: group.createdAt.Format(time.RFC3339),
		UpdatedAt: group.updatedAt.Format(time.RFC3339),
		DeletedAt: helpers.GetPtrValueFromSQLNullTime(group.deletedAt, time.RFC3339),
		Admin: entities.GroupUser{
			ID:        group.admin.id,
			Email:     helpers.GetValueFromSQLNullString(group.admin.email),
			Phone:     helpers.GetValueFromSQLNullString(group.admin.phone),
			FirstName: group.admin.firstName,
			LastName:  helpers.GetPtrValueFromSQLNullString(group.admin.lastName),
			CreatedAt: group.admin.createdAt.Format(time.RFC3339),
			UpdatedAt: group.admin.updatedAt.Format(time.RFC3339),
			DeletedAt: helpers.GetPtrValueFromSQLNullTime(group.admin.deletedAt, time.RFC3339),
			InvitedAt: group.admin.invitedAt.Format(time.RFC3339),
		},
	}, nil
}

func (g *Group) GetUsersList(ctx context.Context, entity entities.GroupUsersListGet) ([]entities.GroupUser, error) {
	query := `
		-- Users groups
		SELECT
   			public.users_groups.invited_at,
   			
   			public.users.id,
   			public.users.email,
   			public.users.phone,
   			public.users.password,
   			public.users.first_name,
   			public.users.last_name,
   			public.users.created_at,
   			public.users.updated_at,
   			public.users.deleted_at
		FROM public.users_groups

		LEFT JOIN public.users
			ON public.users.id = public.users_groups.user_id

		WHERE public.users_groups.group_id = $1
	`

	rows, err := g.postgres.Query(ctx, query, entity.GroupID)
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
			&user.id,
			&user.email,
			&user.phone,
			&user.password,
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
			ID:        user.id,
			Email:     helpers.GetValueFromSQLNullString(user.email),
			Phone:     helpers.GetValueFromSQLNullString(user.phone),
			FirstName: user.firstName,
			LastName:  helpers.GetPtrValueFromSQLNullString(user.lastName),
			CreatedAt: user.createdAt.Format(time.RFC3339),
			UpdatedAt: user.updatedAt.Format(time.RFC3339),
			DeletedAt: helpers.GetPtrValueFromSQLNullTime(user.deletedAt, time.RFC3339),
			InvitedAt: user.invitedAt.Format(time.RFC3339),
		})
	}

	return users, nil
}

func (g *Group) IsGroupAdmin(ctx context.Context, userID, groupID string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM public.users_groups
		WHERE user_id = $1 AND group_id = $2 AND permission = $3
	`

	var count int
	err := g.postgres.QueryRow(ctx, query, userID, groupID, constants.DefaultAdminPermission).Scan(&count)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		logger.Error(ctx, err)
		return false, err
	}
	return count > 0, nil
}
