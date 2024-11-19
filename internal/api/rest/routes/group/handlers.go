package group

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/logger"
	"net/http"
)

func (controller *restGroupController) handlerGroupsList(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}
	logger.Debug(ctx, "Start")

	group, err := controller.useCases.GetList(ctx, entities.GroupsGetList{
		UserID: userID,
	})
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGroupAdd(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	dto := new(AddDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	ctx = logger.WithGroupName(ctx, dto.Name)
	logger.Debug(ctx, "Start")

	err = validateAddDTO(dto)
	if err != nil {
		return err
	}

	group, err := controller.useCases.Add(ctx, entities.GroupAdd{
		Name:    dto.Name,
		AdminID: userID,
	})
	if err != nil {
		return err
	}

	ctx = logger.WithGroupID(ctx, group.ID)
	logger.Debug(ctx, "Finish")
	return c.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGroupGet(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	groupID := c.Param("groupID")
	if groupID == "" {
		return rest_error.ErrInvalidParams
	}

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	logger.Debug(ctx, "Start")

	group, err := controller.useCases.Get(
		ctx,
		userID,
		entities.GroupGet{
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.JSON(http.StatusOK, rest.NewResponseSuccess(group))
}

func (controller *restGroupController) handlerGroupGetInfo(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	groupID := c.Param("groupID")
	if groupID == "" {
		return rest_error.ErrInvalidParams
	}

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	logger.Debug(ctx, "Start")

	groupInfo, err := controller.useCases.GetInfo(
		ctx,
		userID,
		entities.GroupGetInfo{
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.JSON(http.StatusOK, rest.NewResponseSuccess(groupInfo))
}

func (controller *restGroupController) handlerGroupGetUsersList(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	groupID := c.Param("groupID")
	if groupID == "" {
		return rest_error.ErrInvalidParams
	}

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	logger.Debug(ctx, "Start")

	usersList, err := controller.useCases.GetUsersList(
		ctx,
		userID,
		entities.GroupGetUsersList{
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.JSON(http.StatusOK, rest.NewResponseSuccess(usersList))
}

func (controller *restGroupController) handlerGroupPatch(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	groupID := c.Param("groupID")
	if groupID == "" {
		return rest_error.ErrInvalidParams
	}

	userID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	dto := new(PatchDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	if dto.Name != nil {
		ctx = logger.WithGroupName(ctx, *dto.Name)
	}
	logger.Debug(ctx, "Start")

	err = validatePatchDTO(dto)
	if err != nil {
		return err
	}

	err = controller.useCases.Patch(
		ctx,
		userID,
		entities.GroupPatch{
			GroupID: groupID,
			Name:    dto.Name,
		},
	)
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerGroupInviteUser(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	groupID := c.Param("groupID")
	if groupID == "" {
		return rest_error.ErrInvalidParams
	}

	adminID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	dto := new(InviteDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	logger.Debug(ctx, "Start")

	err = validateInviteDTO(dto)
	if err != nil {
		return err
	}
	ctx = logger.With(ctx, logrus.Fields{
		"InviteUserID": dto.UserID,
	})

	err = controller.useCases.InviteUser(
		ctx,
		adminID,
		entities.GroupInviteUser{
			UserID:  dto.UserID,
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}

func (controller *restGroupController) handlerGroupExcludeUser(c echo.Context) error {
	ctx, ok := c.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "No context")
		return rest_error.ErrSomethingHappen
	}

	groupID := c.Param("groupID")
	if groupID == "" {
		return rest_error.ErrInvalidParams
	}

	adminID, ok := c.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	dto := new(ExcludeDTO)
	err := c.Bind(dto)
	if err != nil {
		return rest_error.ErrInvalidBodyData
	}

	logger.Debug(ctx, "Start")

	err = validateExcludeDTO(dto)
	if err != nil {
		return err
	}
	ctx = logger.With(ctx, logrus.Fields{
		"ExcludeUserID": dto.UserID,
	})

	err = controller.useCases.ExcludeUser(
		ctx,
		adminID,
		entities.GroupExcludeUser{
			UserID:  dto.UserID,
			GroupID: groupID,
		},
	)
	if err != nil {
		return err
	}

	logger.Debug(ctx, "Finish")
	return c.NoContent(http.StatusNoContent)
}
