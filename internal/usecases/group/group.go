package group

import (
	"context"
	"fmt"
	"go-clean/config"
	"go-clean/internal/constants/validation_rules"
	"go-clean/internal/entities"
	"go-clean/pkg/helpers"
	"go-clean/pkg/logger"
	"time"
)

func getKeyDeleteGroup(groupID string) string {
	return "group:delete:" + groupID
}

type Group struct {
	cfg          *config.Config
	groupService groupService
	codes        codes
}

func New(cfg *config.Config, groupService groupService, codes codes) Group {
	return Group{
		cfg:          cfg,
		groupService: groupService,
		codes:        codes,
	}
}

func (g *Group) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupService.Add(ctx, entity)
}

func (g *Group) Patch(ctx context.Context, entity entities.GroupPatch) error {
	err := g.groupService.IsGroupAdmin(ctx, entity.UserID, entity.ID)
	if err != nil {
		return err
	}

	err = g.groupService.Patch(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	return g.groupService.GetList(ctx, entity)
}

func (g *Group) Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error) {
	err := g.groupService.IsGroupUser(ctx, entity.UserID, entity.ID)
	if err != nil {
		return nil, err
	}

	return g.groupService.Get(ctx, entity)
}

func (g *Group) Delete(ctx context.Context, entity entities.GroupDelete) error {
	err := g.groupService.IsGroupAdmin(ctx, entity.UserID, entity.ID)
	if err != nil {
		fmt.Println("entity.ID: ", entity.ID)
		fmt.Println("entity.UserID: ", entity.UserID)
		return err
	}

	// Set code to store
	code := helpers.GenerateRandomCode(validation_rules.LenRegistrationCode)
	ctx = logger.WithConfirmationCode(ctx, code)
	logger.Debug(ctx, "Code sent")

	err = g.codes.SetCode(ctx, getKeyDeleteGroup(entity.ID), code, 0, time.Minute*10)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) DeleteConfirm(ctx context.Context, entity entities.GroupDeleteConfirm) error {
	err := g.groupService.IsGroupAdmin(ctx, entity.UserID, entity.ID)
	if err != nil {
		return err
	}

	err = g.codes.CompareCodes(ctx, getKeyDeleteGroup(entity.ID), entity.Code)
	if err != nil {
		return err
	}

	return g.groupService.Delete(ctx, entity)
}
