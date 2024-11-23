package groups

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

type Groups struct {
	cfg           *config.Config
	groupsService groupsService
	codes         codes
}

func New(cfg *config.Config, groupService groupsService, codes codes) Groups {
	return Groups{
		cfg:           cfg,
		groupsService: groupService,
		codes:         codes,
	}
}

func (g *Groups) Add(ctx context.Context, entity entities.GroupAdd) (*entities.Group, error) {
	return g.groupsService.Add(ctx, entity)
}

func (g *Groups) Patch(ctx context.Context, entity entities.GroupPatch) error {
	err := g.groupsService.IsGroupAdmin(ctx, entity.UserID, entity.ID)
	if err != nil {
		return err
	}

	err = g.groupsService.Patch(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}

func (g *Groups) GetList(ctx context.Context, entity entities.GroupsGetList) ([]entities.Group, error) {
	return g.groupsService.GetList(ctx, entity)
}

func (g *Groups) Get(ctx context.Context, entity entities.GroupGetInfo) (*entities.Group, error) {
	err := g.groupsService.IsGroupUser(ctx, entity.UserID, entity.ID)
	if err != nil {
		return nil, err
	}

	return g.groupsService.Get(ctx, entity)
}

func (g *Groups) Delete(ctx context.Context, entity entities.GroupDelete) error {
	err := g.groupsService.IsGroupAdmin(ctx, entity.UserID, entity.ID)
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

func (g *Groups) DeleteConfirm(ctx context.Context, entity entities.GroupDeleteConfirm) error {
	err := g.groupsService.IsGroupAdmin(ctx, entity.UserID, entity.ID)
	if err != nil {
		return err
	}

	err = g.codes.CompareCodes(ctx, getKeyDeleteGroup(entity.ID), entity.Code)
	if err != nil {
		return err
	}

	return g.groupsService.Delete(ctx, entity)
}
