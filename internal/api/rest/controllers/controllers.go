package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/utils/rest_error"
	"go-clean/pkg/logger"
)

// Controller get ctx from echo context

type Handler func(echoCTX echo.Context, ctx context.Context) error

type Controller struct {
	handler Handler
	ctx     context.Context
}

func NewController(handler Handler) *Controller {
	return &Controller{
		handler: handler,
	}
}

func (c *Controller) Register(echoCTX echo.Context) error {
	err := c.Init(echoCTX)
	if err != nil {
		return err
	}

	return c.handler(echoCTX, c.ctx)
}

func (c *Controller) Init(echoCTX echo.Context) error {
	ctx, ok := echoCTX.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "Failed to get context from echo.Context")
		return rest_error.ErrSomethingHappen
	}

	c.ctx = ctx
	return nil
}

// Controller get ctx from echo context and validating dto

type HandlerValidation[V any] func(echoCTX echo.Context, ctx context.Context, dto *V) error

type ControllerValidation[V any] struct {
	handler   HandlerValidation[V]
	validator func(*V) error
	ctx       context.Context
	dto       *V
}

func NewControllerValidation[V any](handler HandlerValidation[V], validator func(*V) error) *ControllerValidation[V] {
	return &ControllerValidation[V]{
		handler:   handler,
		validator: validator,
	}
}

func (c *ControllerValidation[V]) Register(echoCTX echo.Context) error {
	err := c.Init(echoCTX)
	if err != nil {
		return rest_error.ErrSomethingHappen
	}

	return c.handler(echoCTX, c.ctx, c.dto)
}

func (c *ControllerValidation[V]) Init(echoCTX echo.Context) error {
	ctx, ok := echoCTX.Get(constants.CTXLoggerInCTX).(context.Context)
	if !ok {
		logger.Error(ctx, "Failed to get context from echo.Context")
		return rest_error.ErrSomethingHappen
	}

	dto := new(V)
	err := echoCTX.Bind(dto)
	if err != nil {
		logger.Debug(ctx, "Failed to bind request body", "error", err)
		return rest_error.ErrInvalidBodyData
	}

	err = c.validator(dto)
	if err != nil {
		logger.Debug(ctx, "Validation failed", "error", err)
		return err
	}

	c.ctx = ctx
	c.dto = dto
	return nil
}
