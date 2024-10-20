package controller

import (
	"easy-life-back-go/internal/constants/validation_rules"
	"easy-life-back-go/internal/server/routes/auth/service"
	"easy-life-back-go/internal/server/utils/response"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service *service.Service
	logger  echo.Logger
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
		return ctx.JSON(http.StatusBadRequest, response.NewBad(response.InvalidJSON))
	}

	err = validation.ValidateStruct(signInData,
		validation.Field(&signInData.Name, validation.Required, validation.RuneLength(validation_rules.LenMinName, validation_rules.LenMaxName)),
		validation.Field(&signInData.Email, validation.Required, is.Email),
		validation.Field(&signInData.Password, validation.Required, validation.RuneLength(validation_rules.LenMinPassword, validation_rules.LenMaxPassword)),
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.NewBadValidation(err.Error()))
	}

	return nil
}

func (controller *Controller) Registration(ctx echo.Context) error {
	return nil
}

func (controller *Controller) RegistrationSuccess(ctx echo.Context) error {
	return nil
}
