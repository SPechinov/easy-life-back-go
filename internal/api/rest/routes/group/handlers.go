package group

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/utils"
	"go-clean/internal/entities"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restGroupController) handlerGetGroupsList(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}

	group, err := controller.useCases.GetList(
		ctx,
		entities.GroupsGetList{UserID: userID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerAddGroup(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}

	dto, err := utils.BindAndValidate[AddDTO](echoCtx, validateAddDTO)
	if err != nil {
		return err
	}

	ctx = logger.WithGroupName(ctx, dto.Name)

	group, err := controller.useCases.Add(
		ctx,
		entities.GroupAdd{
			Name:    dto.Name,
			AdminID: userID,
		},
	)
	if err != nil {
		return err
	}

	ctx = logger.WithGroupID(ctx, group.ID)
	logger.Info(ctx, "Group created")
	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGetGroup(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}

	groupID := echoCtx.Param("groupID")

	group, err := controller.useCases.Get(
		ctx,
		userID,
		entities.GroupGet{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGetGroupInfo(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}

	groupID := echoCtx.Param("groupID")

	groupInfo, err := controller.useCases.GetInfo(
		ctx,
		userID,
		entities.GroupGetInfo{ID: groupID},
	)
	if err != nil {
		return err
	}

	return echoCtx.JSON(http.StatusOK, rest.NewResponseSuccess(groupInfo))
}

func (controller *restGroupController) handlerGetGroupUsers(echoCtx echo.Context) error {
	ctx, err := utils.GetCTXLoggerFromEchoCTX(echoCtx)
	if err != nil {
		return err
	}
	userID, err := utils.GetUserIDFromEchoCTX(echoCtx)
	if err != nil {
		return err
	}
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

func (controller *restGroupController) handlerPatchGroup(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}

	groupID := echoCtx.Param("groupID")

	dto, err := utils.BindAndValidate[PatchDTO](echoCtx, validatePatchDTO)
	if err != nil {
		return err
	}

	err = controller.useCases.Patch(
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

func (controller *restGroupController) handlerInviteUserInGroup(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}
	groupID := echoCtx.Param("groupID")

	dto, err := utils.BindAndValidate[InviteUserDTO](echoCtx, validateInviteUserDTO)
	if err != nil {
		return err
	}

	err = controller.useCases.InviteUser(
		ctx,
		userID,
		entities.GroupInviteUser{
			UserID: dto.UserID,
			ID:     groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Info(ctx, fmt.Sprintf("User invited: %s", dto.UserID))
	return echoCtx.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerExcludeUserFromGroup(echoCtx echo.Context) error {
	ctx, userID, err := initializeRequest(echoCtx)
	if err != nil {
		return err
	}
	groupID := echoCtx.Param("groupID")

	dto, err := utils.BindAndValidate[ExcludeUserDTO](echoCtx, validateExcludeUserDTO)
	if err != nil {
		return err
	}

	err = controller.useCases.ExcludeUser(
		ctx,
		userID,
		entities.GroupExcludeUser{
			UserID: dto.UserID,
			ID:     groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Info(ctx, fmt.Sprintf("User excluded: %s", dto.UserID))
	return echoCtx.NoContent(http.StatusNoContent)
}

func initializeRequest(echoCtx echo.Context) (context.Context, string, error) {
	ctx, err := utils.GetCTXLoggerFromEchoCTX(echoCtx)
	if err != nil {
		return nil, "", err
	}
	userID, err := utils.GetUserIDFromEchoCTX(echoCtx)
	if err != nil {
		return nil, "", err
	}

	return ctx, userID, nil
}
