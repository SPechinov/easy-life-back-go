package controller

import (
	"easy-life-back-go/internal/server/routes/auth/service"
	"easy-life-back-go/internal/server/routes/auth/validation"
	"easy-life-back-go/internal/server/routes/auth/views"
	"easy-life-back-go/internal/server/utils/response"
	"errors"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Controller struct {
	service   *service.Service
	validator *validation.Validator
}

func NewController(
	service *service.Service,
	validator *validation.Validator,
) *Controller {
	return &Controller{
		service:   service,
		validator: validator,
	}
}

func (c *Controller) SignIn(ctx echo.Context) error {
	data := new(views.SignInData)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.CodeInvalidJSON))
	}

	err = c.validator.SignIn(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	return nil
}

func (c *Controller) Registration(ctx echo.Context) error {
	data := new(views.RegistrationData)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.CodeInvalidJSON))
	}

	err = c.validator.Registration(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	err = c.service.Registration(data.Email)
	if err != nil {
		return c.handleServiceError(ctx, err, "auth_controller_registration_success")
	}

	return ctx.JSON(http.StatusOK, response.NewSuccess(nil))
}

func (c *Controller) RegistrationSuccess(ctx echo.Context) error {
	data := new(views.RegistrationSuccessData)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.CodeInvalidJSON))
	}

	err = c.validator.RegistrationSuccess(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	err = c.service.RegistrationSuccess(data.Name, data.Email, data.Password, data.Code)
	if err != nil {
		return c.handleServiceError(ctx, err, "auth_controller_registration_success")
	}

	return ctx.JSON(http.StatusOK, response.NewSuccess(nil))
}

func (c *Controller) handleServiceError(ctx echo.Context, err error, logMessage string) error {
	var controllerErr *response.Bad
	if errors.As(err, &controllerErr) {
		return ctx.JSON(controllerErr.HttpCode, controllerErr.Info)
	}
	slog.Error(logMessage, "err_data", err)
	return ctx.JSON(http.StatusInternalServerError, response.NewBadInfo(response.CodeSomethingHappen))
}
