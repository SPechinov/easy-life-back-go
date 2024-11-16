package group

import (
	"context"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
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

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
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
		RETURNING id, name, is_payed, created_at, updated_at, deleted_at
	`

	group := new(dataGroup)
	err = tx.QueryRow(ctx, queryAddGroup, entity.Name).Scan(
		&group.ID,
		&group.Name,
		&group.IsPayed,
		&group.CreatedAt,
		&group.UpdatedAt,
		&group.DeletedAt,
	)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	queryUsersGroup := `
		INSERT INTO public.users_groups (group_id, user_id, permission)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(ctx, queryUsersGroup, group.ID, entity.AdminID, constants.DefaultAdminPermission)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &entities.Group{
		ID:        group.ID,
		Name:      group.Name,
		CreatedAt: group.CreatedAt.Format(time.RFC3339),
		UpdatedAt: group.UpdatedAt.Format(time.RFC3339),
	}, nil
}
