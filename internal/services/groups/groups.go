package groups

import (
	"context"
	"go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/client_error"
)

type Groups struct {
	groupsDatabase groupsDatabase
}

func New(groupDatabase groupsDatabase) *Groups {
	return &Groups{
		groupsDatabase: groupDatabase,
	}
}

func (g *Groups) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupsDatabase.Add(ctx, entity)
}

func (g *Groups) Patch(ctx context.Context, entity entities.GroupPatch) error {
	return g.groupsDatabase.Patch(ctx, entity)
}

func (g *Groups) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	return g.groupsDatabase.GetList(ctx, entity)
}

func (g *Groups) Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error) {
	return g.groupsDatabase.Get(ctx, entities.GroupGet{ID: entity.ID})
}

func (g *Groups) Delete(ctx context.Context, entity entities.GroupDeleteConfirm) error {
	return g.groupsDatabase.Delete(ctx, entity)
}

func (g *Groups) GetGroupUser(ctx context.Context, userID, groupID string) (*entities.GroupUser, error) {
	return g.groupsDatabase.GetGroupUser(ctx, userID, groupID)
}

func (g *Groups) IsGroupAdmin(ctx context.Context, userID, groupID string) error {
	user, err := g.groupsDatabase.GetGroupUser(ctx, userID, groupID)
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

func (g *Groups) IsGroupUser(ctx context.Context, userID, groupID string) error {
	user, err := g.GetGroupUser(ctx, userID, groupID)
	if user == nil && err == nil {
		return client_error.ErrUserNotInGroup
	}
	if err != nil {
		return err
	}

	return nil
}
