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

func (controller *restGroupController) handlerGetFullGroup(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")

	group, err := controller.useCases.GetFull(
		ctx,
		userID,
		entities.GroupGet{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGetGroupInfo(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")

	groupInfo, err := controller.useCases.Get(
		ctx,
		userID,
		entities.GroupGetInfo{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, rest.NewResponseSuccess(groupInfo))
}

func (controller *restGroupController) handlerGetGroupUsers(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")

	usersList, err := controller.useCases.GetUsersList(
		ctx,
		userID,
		entities.GroupGetUsersList{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCTX.JSON(http.StatusOK, rest.NewResponseSuccess(usersList))
}

func (controller *restGroupController) handlerPatchGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *PatchDTO) error {
	groupID := echoCTX.Param("groupID")
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
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerInviteUserInGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *InviteUserDTO) error {
	groupID := echoCTX.Param("groupID")
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
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerExcludeUserFromGroup(echoCTX echo.Context, ctx context.Context, userID string, dto *ExcludeUserDTO) error {
	groupID := echoCTX.Param("groupID")
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
	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerDelete(echoCTX echo.Context, ctx context.Context, userID string) error {
	groupID := echoCTX.Param("groupID")

	err := controller.useCases.Delete(ctx, groupID, userID)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerDeleteConfirm(echoCTX echo.Context, ctx context.Context, userID string, dto *DeleteDTO) error {
	groupID := echoCTX.Param("groupID")

	err := controller.useCases.DeleteConfirm(ctx, groupID, userID, dto.Code)
	if err != nil {
		return err
	}

	return echoCTX.NoContent(http.StatusNoContent)
}
