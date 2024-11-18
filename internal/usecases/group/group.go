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

func (g *Group) GetUsersList(ctx context.Context, userID string, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	user, err := g.groupService.GetGroupUser(ctx, userID, entity.GroupID)
	if user == nil && err == nil {
		return nil, client_error.ErrUserNotInGroup
	}

	users, err := g.groupService.GetUsersList(ctx, entity)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (g *Group) InviteUser(ctx context.Context, adminID string, entity entities.GroupInviteUser) error {
	isAdmin, err := g.groupService.IsGroupAdmin(ctx, adminID, entity.GroupID)
	if err != nil {
		return err
	}
	if !isAdmin {
		return client_error.ErrUserNotAdminGroup
	}

	err = g.groupService.InviteUser(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) ExcludeUser(ctx context.Context, adminID string, entity entities.GroupExcludeUser) error {
	isAdmin, err := g.groupService.IsGroupAdmin(ctx, adminID, entity.GroupID)
	if err != nil {
		return err
	}
	if !isAdmin {
		return client_error.ErrUserNotAdminGroup
	}
	if adminID == entity.UserID {
		return client_error.ErrUserAdminGroup
	}

	err = g.groupService.ExcludeUser(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
