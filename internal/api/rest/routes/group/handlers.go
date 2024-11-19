package group

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restGroupController) handlerGetGroupsList(echoCtx echo.Context) error {
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	userID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)

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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	userID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)

	dto := new(AddDTO)
	err := echoCtx.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithGroupName(ctx, dto.Name)

	err = validateAddDTO(dto)
	if err != nil {
		return err
	}

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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	userID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)
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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	userID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)
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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	userID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)
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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	groupID := echoCtx.Param("groupID")
	userID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)

	dto := new(PatchDTO)
	err := echoCtx.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	if dto.Name != nil {
		ctx = logger.WithGroupName(ctx, *dto.Name)
	}

	err = validatePatchDTO(dto)
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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	adminID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)
	groupID := echoCtx.Param("groupID")

	dto := new(InviteUserDTO)
	err := echoCtx.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	err = validateInviteDTO(dto)
	if err != nil {
		return err
	}

	err = controller.useCases.InviteUser(
		ctx,
		adminID,
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
	ctx := echoCtx.Get(constants.CTXLoggerInCTX).(context.Context)
	adminID := echoCtx.Get(globalConstants.CTXUserIDKey).(string)
	groupID := echoCtx.Param("groupID")

	dto := new(ExcludeUserDTO)
	err := echoCtx.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	err = validateExcludeDTO(dto)
	if err != nil {
		return err
	}

	err = controller.useCases.ExcludeUser(
		ctx,
		adminID,
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
