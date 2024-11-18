package group

import (
	"context"
	"go-clean/config"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
)

type Group struct {
	cfg          *config.Config
	groupService groupService
}

func New(cfg *config.Config, groupService groupService) Group {
	return Group{
		cfg:          cfg,
		groupService: groupService,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupService.Add(ctx, entity)
}

func (g *Group) Patch(ctx context.Context, adminID string, entity entities.GroupPatch) (*entities.Group, error) {
	isAdmin, err := g.groupService.IsGroupAdmin(ctx, adminID, entity.GroupID)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, client_error.ErrUserNotAdminGroup
	}
	err = g.groupService.Patch(ctx, entity)
	if err != nil {
		return nil, err
	}

	group, err := g.groupService.Get(ctx, entities.GroupGet{GroupID: entity.GroupID})
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (g *Group) Get(ctx context.Context, userID string, entity entities.GroupGet) (*entities.Group, error) {
	user, err := g.groupService.GetGroupUser(ctx, userID, entity.GroupID)
	if user == nil && err == nil {
		return nil, client_error.ErrUserNotInGroup
	}
	if err != nil {
		return nil, err
	}

	return g.groupService.Get(ctx, entity)
}
