package group

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest"
	"go-clean/internal/entities"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restGroupController) handlerGetGroupsList(echoCtx echo.Context, ctx context.Context, userID string) error {
	group, err := controller.useCases.GetList(
		ctx,
		entities.GroupsGetList{UserID: userID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerAddGroup(echoCtx echo.Context, ctx context.Context, dto *AddDTO, userID string) error {
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
	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGetFullGroup(echoCtx echo.Context, ctx context.Context, userID string) error {
	groupID := echoCtx.Param("groupID")

	group, err := controller.useCases.GetFull(
		ctx,
		userID,
		entities.GroupGet{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGetGroupInfo(echoCtx echo.Context, ctx context.Context, userID string) error {
	groupID := echoCtx.Param("groupID")

	groupInfo, err := controller.useCases.Get(
		ctx,
		userID,
		entities.GroupGetInfo{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(groupInfo))
}

func (controller *restGroupController) handlerGetGroupUsers(echoCtx echo.Context, ctx context.Context, userID string) error {
	groupID := echoCtx.Param("groupID")

	usersList, err := controller.useCases.GetUsersList(
		ctx,
		userID,
		entities.GroupGetUsersList{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(usersList))
}

func (controller *restGroupController) handlerPatchGroup(echoCtx echo.Context, ctx context.Context, dto *PatchDTO, userID string) error {
	groupID := echoCtx.Param("groupID")
	err := controller.useCases.Patch(
		ctx,
		userID,
		entities.GroupPatch{
			ID:   groupID,
			Name: dto.Name,
		},
	)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Group updated")
	return echoCtx.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerInviteUserInGroup(echoCtx echo.Context, ctx context.Context, dto *InviteUserDTO, userID string) error {
	groupID := echoCtx.Param("groupID")
	err := controller.useCases.InviteUser(
		ctx,
		userID,
		entities.GroupInviteUser{
			ID:     groupID,
			UserID: dto.UserID,
		},
	)
	if err != nil {
		return err
	}

	logger.Info(ctx, fmt.Sprintf("User invited: %s", dto.UserID))
	return echoCtx.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerExcludeUserFromGroup(echoCtx echo.Context, ctx context.Context, dto *ExcludeUserDTO, userID string) error {
	groupID := echoCtx.Param("groupID")
	err := controller.useCases.ExcludeUser(
		ctx,
		userID,
		entities.GroupExcludeUser{
			ID:     groupID,
			UserID: dto.UserID,
		},
	)
	if err != nil {
		return err
	}

	logger.Info(ctx, fmt.Sprintf("User excluded: %s", dto.UserID))
	return echoCtx.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerDelete(echoCtx echo.Context, ctx context.Context, userID string) error {
	groupID := echoCtx.Param("groupID")

	err := controller.useCases.Delete(ctx, groupID, userID)
	if err != nil {
		return err
	}

	return echoCtx.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerDeleteConfirm(echoCtx echo.Context, ctx context.Context, dto *DeleteDTO, userID string) error {
	groupID := echoCtx.Param("groupID")

	err := controller.useCases.DeleteConfirm(ctx, groupID, userID, dto.Code)
	if err != nil {
		return err
	}

	return echoCtx.NoContent(http.StatusNoContent)
}
