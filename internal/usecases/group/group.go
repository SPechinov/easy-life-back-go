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

func (g *Group) Patch(ctx context.Context, adminID string, entity entities.GroupPatch) error {
	isDeletedGroup := g.groupService.IsDeletedGroup(ctx, entity.ID)
	if isDeletedGroup {
		return client_error.ErrGroupDeleted
	}

	isAdmin := g.groupService.IsGroupAdmin(ctx, adminID, entity.ID)
	if !isAdmin {
		return client_error.ErrUserNotAdminGroup
	}

	err := g.groupService.Patch(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) GetFull(ctx context.Context, userID string, entity entities.GroupGet) (*entities.GroupFull, error) {
	user, err := g.groupService.GetGroupUser(ctx, userID, entity.ID)
	if user == nil && err == nil {
		return nil, client_error.ErrUserNotInGroup
	}
	if err != nil {
		return nil, err
	}

	groupFull, err := g.groupService.GetFull(ctx, entity)
	if err != nil {
		return nil, err
	}

	if groupFull.Deleted() {
		return nil, client_error.ErrGroupDeleted
	}

	return groupFull, nil
}

func (g *Group) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	return g.groupService.GetList(ctx, entity)
}

func (g *Group) Get(ctx context.Context, userID string, entity entities.GroupGetInfo) (*entities.Group, error) {
	user, err := g.groupService.GetGroupUser(ctx, userID, entity.ID)
	if user == nil && err == nil {
		return nil, client_error.ErrUserNotInGroup
	}
	if err != nil {
		return nil, err
	}

	group, err := g.groupService.Get(ctx, entity)
	if err != nil {
		return nil, err
	}

	if group.Deleted() {
		return nil, client_error.ErrGroupDeleted
	}

	return group, nil
}

func (g *Group) GetUsersList(ctx context.Context, userID string, entity entities.GroupGetUsersList) ([]entities.GroupUser, error) {
	isDeletedGroup := g.groupService.IsDeletedGroup(ctx, entity.ID)
	if isDeletedGroup {
		return nil, client_error.ErrGroupDeleted
	}

	user, err := g.groupService.GetGroupUser(ctx, userID, entity.ID)
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
	isDeletedGroup := g.groupService.IsDeletedGroup(ctx, entity.ID)
	if isDeletedGroup {
		return client_error.ErrGroupDeleted
	}

	isAdmin := g.groupService.IsGroupAdmin(ctx, adminID, entity.ID)
	if !isAdmin {
		return client_error.ErrUserNotAdminGroup
	}

	err := g.groupService.InviteUser(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) ExcludeUser(ctx context.Context, adminID string, entity entities.GroupExcludeUser) error {
	isDeletedGroup := g.groupService.IsDeletedGroup(ctx, entity.ID)
	if isDeletedGroup {
		return client_error.ErrGroupDeleted
	}

	isAdmin := g.groupService.IsGroupAdmin(ctx, adminID, entity.ID)
	if !isAdmin {
		return client_error.ErrUserNotAdminGroup
	}
	if adminID == entity.UserID {
		return client_error.ErrUserAdminGroup
	}

	err := g.groupService.ExcludeUser(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
