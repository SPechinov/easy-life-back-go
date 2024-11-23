package group

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest"
	"go-clean/internal/entities"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restGroupController) handlerGetGroupsList(echoCTX echo.Context, ctx context.Context, userID string) error {
	group, err := controller.useCases.GetList(
		ctx,
		entities.GroupsGetList{UserID: userID},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerAddGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *AddDTO) error {
	ctx = logger.WithGroupName(ctx, dto.Name)

	group, err := controller.useCases.Add(
		ctx,
		entities.GroupAdd{
			AdminID: userID,
			Name:    dto.Name,
		},
	)
	if err != nil {
		return err
	}

	ctx = logger.WithGroupID(ctx, group.ID)
	logger.Info(ctx, "Group created")
	return echoCTX.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGetGroup(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")

	groupInfo, err := controller.useCases.Get(
		ctx,
		entities.GroupGetInfo{
			ID:     groupID,
			UserID: userID,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, rest.NewResponseSuccess(groupInfo))
}

func (controller *restGroupController) handlerPatchGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *PatchDTO) error {
	groupID := echoCTX.Param("groupID")
	err := controller.useCases.Patch(
		ctx,
		entities.GroupPatch{
			ID:     groupID,
			UserID: userID,
			Name:   dto.Name,
		},
	)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Group updated")
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerDelete(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")

	err := controller.useCases.Delete(
		ctx,
		entities.GroupDelete{
			ID:     userID,
			UserID: groupID,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerDeleteConfirm(echoCTX echo.Context, ctx context.Context, userID string, dto *DeleteConfirmDTO) error {
	groupID := echoCTX.Param("groupID")

	err := controller.useCases.DeleteConfirm(
		ctx,
		entities.GroupDeleteConfirm{
			ID:     groupID,
			UserID: userID,
			Code:   dto.Code,
		},
	)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}
