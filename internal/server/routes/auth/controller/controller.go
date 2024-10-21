package controller

import (
	"easy-life-back-go/internal/constants/validation_rules"
	"easy-life-back-go/internal/server/routes/auth/service"
	"easy-life-back-go/internal/server/utils/response"
	"errors"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) SignIn(ctx echo.Context) error {
	data := new(SignInData)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.CodeInvalidJSON))
	}

	err = validation.ValidateStruct(data,
		validation.Field(&data.Email, validation.Required, is.Email),
		validation.Field(&data.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	return nil
}

func (controller *Controller) Registration(ctx echo.Context) error {
	data := new(RegistrationData)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.CodeInvalidJSON))
	}

	err = validation.ValidateStruct(data,
		validation.Field(&data.Email, validation.Required, is.Email),
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	err = controller.service.Registration(data.Email)

	if err != nil {
		var controllerErr *response.Bad
		if errors.As(err, &controllerErr) {
			return ctx.JSON(controllerErr.HttpCode, controllerErr.Info)
		}

		slog.Error("auth_controller_registration", "err_data", err)
		return ctx.JSON(http.StatusInternalServerError, response.NewBadInfo(response.CodeSomethingHappen))
	}

	return ctx.JSON(http.StatusOK, response.NewSuccess(nil))
}

func (controller *Controller) RegistrationSuccess(ctx echo.Context) error {
	data := new(RegistrationSuccessData)

	err := ctx.Bind(data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.CodeInvalidJSON))
	}

	err = validation.ValidateStruct(data,
		validation.Field(&data.Email, validation.Required, is.Email),
		validation.Field(&data.Name, validation.Required, validation.RuneLength(validation_rules.LenMinName, validation_rules.LenMaxName)),
		validation.Field(&data.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
		validation.Field(&data.Code, validation.Required, validation.RuneLength(validation_rules.LenRegistrationCode, validation_rules.LenRegistrationCode)),
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	err = controller.service.RegistrationSuccess(data.Name, data.Email, data.Password, data.Code)

	if err != nil {
		var controllerErr *response.Bad
		if errors.As(err, &controllerErr) {
			return ctx.JSON(controllerErr.HttpCode, controllerErr.Info)
		}

		slog.Error("auth_controller_registration_success", "err_data", err)
		return ctx.JSON(http.StatusInternalServerError, response.NewBadInfo(response.CodeSomethingHappen))
	}

	return ctx.JSON(http.StatusOK, response.NewSuccess(nil))
}
