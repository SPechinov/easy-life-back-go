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
		SET 
		    name = COALESCE(NULLIF($1, ''), name),
			deleted_at = CASE WHEN $2 = true THEN CURRENT_TIMESTAMP ELSE deleted_at END
		WHERE id = $3
	`

	_, err := g.postgres.Exec(ctx, query, entity.Name, entity.Delete, entity.ID)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

func (g *Group) Get(ctx context.Context, entity entities.GroupGet) (*entities.Group, error) {
	query := `
		SELECT
			public.groups.id AS group_id,
			public.groups.name AS group_name,
			public.groups.is_payed AS group_is_payed,
			public.groups.created_at AS group_created_at,
			public.groups.updated_at AS group_updated_at,
			public.groups.deleted_at AS group_deleted_at
		FROM public.groups WHERE public.groups.id = $1
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

func (g *Group) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	query := `
		SELECT
		    public.groups.id,
		    public.groups.name,
		    public.groups.is_payed,
		    public.groups.created_at,
		    public.groups.updated_at,
		    public.groups.deleted_at
		FROM public.users_groups

		LEFT JOIN public.groups
			ON public.groups.id = public.users_groups.group_id

		WHERE 
		    user_id = $1 
		  AND (
		    ($2 = TRUE AND public.groups.deleted_at IS NOT NULL)
			OR
		    ($2 = FALSE AND public.groups.deleted_at IS NULL)
		  )
	`

	rows, err := g.postgres.Query(ctx, query, entity.UserID, entity.Deleted)
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

func (g *Group) GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error) {
	query := `
		-- Users groups
		SELECT
   			public.users_groups.invited_at,
   			public.users_groups.permission,
   			
   			public.users.id,
   			public.users.email,
   			public.users.phone,
   			public.users.first_name,
   			public.users.last_name,
   			public.users.created_at,
   			public.users.updated_at,
   			public.users.deleted_at
		FROM public.users_groups

		LEFT JOIN public.users
			ON public.users.id = public.users_groups.user_id

		WHERE  public.users_groups.user_id = $1 AND public.users_groups.group_id = $2
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
