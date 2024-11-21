package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
	"go-clean/pkg/logger"
)

// Controller get ctx and UserID from echo context and

type HandlerUserID func(echoCTX echo.Context, ctx context.Context, userID string) error

type ControllerUserID struct {
	Controller
	handler HandlerUserID
	userID  string
}

func NewControllerUserID(handler HandlerUserID) *ControllerUserID {
	return &ControllerUserID{
		handler: handler,
	}
}

func (c *ControllerUserID) Register(echoCTX echo.Context) error {
	err := c.Init(echoCTX)
	if err != nil {
		return err
	}

	c.ctx = logger.WithUserID(c.ctx, c.userID)

	return c.handler(echoCTX, c.ctx, c.userID)
}

func (c *ControllerUserID) Init(echoCTX echo.Context) error {
	err := c.Controller.Init(echoCTX)

	if err != nil {
		return err
	}

	userID, ok := echoCTX.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	c.userID = userID
	return nil
}

type HandlerUserIDValidation[V any] func(echoCTX echo.Context, ctx context.Context, userID string, dto *V) error

type ControllerUserIDValidation[V any] struct {
	ControllerValidation[V]
	handler HandlerUserIDValidation[V]
	userID  string
}

func NewControllerUserIDValidation[V any](handler HandlerUserIDValidation[V], validator func(*V) error) *ControllerUserIDValidation[V] {
	return &ControllerUserIDValidation[V]{
		ControllerValidation: ControllerValidation[V]{
			validator: validator,
		},
		handler: handler,
	}
}

func (c *ControllerUserIDValidation[V]) Register(echoCTX echo.Context) error {
	err := c.Init(echoCTX)
	if err != nil {
		return err
	}

	c.ctx = logger.WithUserID(c.ctx, c.userID)

	return c.handler(echoCTX, c.ctx, c.userID, c.dto)
}

func (c *ControllerUserIDValidation[V]) Init(echoCTX echo.Context) error {
	err := c.ControllerValidation.Init(echoCTX)

	if err != nil {
		return err
	}

	userID, ok := echoCTX.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	c.userID = userID
	return nil
}
