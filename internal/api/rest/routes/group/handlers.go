package group

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
	"go-clean/internal/entities"
	"go-clean/pkg/logger"
	"net/http"
)

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
