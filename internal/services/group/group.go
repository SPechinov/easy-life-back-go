package group

import (
	"context"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
)

type Group struct {
	groupDatabase groupDatabase
}

func New(groupDatabase groupDatabase) *Group {
	return &Group{
		groupDatabase: groupDatabase,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupDatabase.Add(ctx, entity)
}

func (g *Group) Patch(ctx context.Context, entity entities.GroupPatch) error {
	return g.groupDatabase.Patch(ctx, entity)
}

func (g *Group) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	return g.groupDatabase.GetList(ctx, entity)
}

func (g *Group) Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error) {
	return g.groupDatabase.Get(ctx, entities.GroupGet{ID: entity.ID})
}

func (g *Group) Delete(ctx context.Context, entity entities.GroupDeleteConfirm) error {
	return g.groupDatabase.Delete(ctx, entity)
}

func (g *Group) GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error) {
	return g.groupDatabase.GetGroupUser(ctx, userID, groupID)
}

func (g *Group) IsGroupAdmin(ctx context.Context, userID, groupID string) error {
	user, err := g.groupDatabase.GetGroupUser(ctx, userID, groupID)
	if user == nil && err == nil {
		return client_error.ErrUserNotInGroup
	}
	if err != nil {
		return err
	}

	if user.Permission != constants.DefaultAdminPermission {
		return client_error.ErrUserNotAdminGroup
	}

	return nil
}

func (g *Group) IsGroupUser(ctx context.Context, userID, groupID string) error {
	user, err := g.GetGroupUser(ctx, userID, groupID)
	if user == nil && err == nil {
		return client_error.ErrUserNotInGroup
	}
	if err != nil {
		return err
	}

	return nil
}
