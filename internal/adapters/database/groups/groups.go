package groups

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"go-clean/pkg/postgres"
	"time"
)

type Groups struct {
	postgres postgres.Client
}

func New(postgres postgres.Client) *Groups {
	return &Groups{
		postgres: postgres,
	}
}

func (g *Groups) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	tx, err := g.postgres.Begin(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	queryAddGroup := `
		INSERT INTO public.groups (name)
		VALUES ($1)
		RETURNING id, name, created_at, updated_at
	`

	group := new(dataGroup)
	err = tx.QueryRow(ctx, queryAddGroup, entity.Name).Scan(
		&group.id,
		&group.name,
		&group.createdAt,
		&group.updatedAt,
	)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	queryUsersGroup := `
		INSERT INTO public.groups_users (group_id, user_id, permission)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(ctx, queryUsersGroup, group.id, entity.AdminID, constants.DefaultAdminPermission)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &entities.Group{
		ID:        group.id,
		Name:      group.name,
		IsPayed:   false,
		CreatedAt: group.createdAt.Format(time.RFC3339),
		UpdatedAt: group.updatedAt.Format(time.RFC3339),
		DeletedAt: nil,
	}, nil
}

func (g *Groups) Patch(ctx context.Context, entity entities.GroupPatch) error {
	query := `
		UPDATE public.groups
		SET name = COALESCE(NULLIF($1, ''), name)
		WHERE id = $3 AND deleted_at IS NULL
	`

	_, err := g.postgres.Exec(ctx, query, entity.Name, entity.ID)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (g *Groups) Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error) {
	query := `
		SELECT id, name, is_payed, created_at, updated_at, deleted_at
		FROM public.groups
		WHERE id = $1 AND deleted_at IS NULL
	`

	group := new(dataGroup)
	err := g.postgres.QueryRow(ctx, query, entity.ID).Scan(
		&group.id,
		&group.name,
		&group.isPayed,
		&group.createdAt,
		&group.updatedAt,
		&group.deletedAt,
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
	}, nil
}

func (g *Groups) Delete(ctx context.Context, entity entities.GroupDeleteConfirm) error {
	query := `
		UPDATE public.groups
		SET deleted_at = NOW()
	   	WHERE groups.id = $1 AND deleted_at IS NULL
	`

	res, err := g.postgres.Exec(ctx, query, entity.ID)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if res.RowsAffected() == 0 {
		return client_error.ErrGroupDeleted
	}

	return nil
}

func (g *Groups) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	query := `
		SELECT
		    public.groups.id,
		    public.groups.name,
		    public.groups.is_payed,
		    public.groups.created_at,
		    public.groups.updated_at,
		    public.groups.deleted_at
		FROM public.groups_users
		LEFT JOIN public.groups ON public.groups.id = public.groups_users.group_id
		WHERE public.groups_users.user_id = $1 AND public.groups.deleted_at IS NULL
	`

	rows, err := g.postgres.Query(ctx, query, entity.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	defer rows.Close()

	groups := make([]entities.Group, 0)

	for rows.Next() {
		var group dataGroup
		err = rows.Scan(
			&group.id,
			&group.name,
			&group.isPayed,
			&group.createdAt,
			&group.updatedAt,
			&group.deletedAt,
		)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		groups = append(groups, entities.Group{
			ID:        group.id,
			Name:      group.name,
			IsPayed:   group.isPayed,
			CreatedAt: group.createdAt.Format(time.RFC3339),
			UpdatedAt: group.updatedAt.Format(time.RFC3339),
			DeletedAt: helpers.GetPtrValueFromSQLNullTime(group.deletedAt, time.RFC3339),
		})
	}

	return groups, nil
}

func (g *Groups) GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error) {
	query := `
		-- Users groups
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
		LEFT JOIN public.groups ON groups.id = public.groups_users.group_id
		WHERE 
			public.groups_users.user_id = $1
	  		AND public.groups_users.group_id = $2 
	  		AND public.groups.deleted_at IS NULL
	`

	var user dataUser
	err := g.postgres.QueryRow(ctx, query, userID, groupID).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		logger.Error(ctx, err)
		return nil, err
	}

	return &entities.GroupUser{
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
	}, nil
}
