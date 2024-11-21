package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-clean/internal/api/rest/utils/rest_error"
	globalConstants "go-clean/internal/constants"
)

// Controller get ctx and UserID from echo context and

type HandlerUserID func(echoCTX echo.Context, ctx context.Context, userID string) error

type ControllerUserID struct {
	Controller
	handler HandlerUserID
	userID  string
}

func NewControllerWithUserID(handler HandlerUserID) *ControllerUserID {
	return &ControllerUserID{
		handler: handler,
	}
}

func (c *ControllerUserID) Register(echoCTX echo.Context) error {
	err := c.Controller.Init(echoCTX)
	if err != nil {
		return err
	}

	userID, ok := echoCTX.Get(globalConstants.CTXUserIDKey).(string)
	if !ok {
		return rest_error.ErrNotAuthorized
	}

	return c.handler(echoCTX, c.ctx, userID)
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
