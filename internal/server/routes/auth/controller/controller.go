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
	signInData := new(SignInData)

	err := ctx.Bind(signInData)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.InvalidJSON))
	}

	err = validation.ValidateStruct(signInData,
		validation.Field(&signInData.Email, validation.Required, is.Email),
		validation.Field(&signInData.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	return nil
}

func (controller *Controller) Registration(ctx echo.Context) error {
	registrationData := new(RegistrationData)

	err := ctx.Bind(registrationData)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.InvalidJSON))
	}

	err = validation.ValidateStruct(registrationData,
		validation.Field(&registrationData.Name, validation.Required, validation.RuneLength(validation_rules.LenMinName, validation_rules.LenMaxName)),
		validation.Field(&registrationData.Email, validation.Required, is.Email),
		validation.Field(&registrationData.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfoValidation(err.Error()))
	}

	err = controller.service.Registration(registrationData.Name, registrationData.Email, registrationData.Password)

	if err != nil {
		var controllerError *response.Bad
		if errors.As(err, &controllerError) {
			return ctx.JSON(controllerError.HttpCode, controllerError.Info)
		}

		slog.Error("auth_controller_registration", "err_data", err)
		return ctx.JSON(http.StatusBadRequest, response.NewBadInfo(response.SomethingHappen))
	}

	return ctx.JSON(http.StatusOK, response.NewSuccess(nil))
}

func (controller *Controller) RegistrationSuccess(ctx echo.Context) error {
	return nil
}
