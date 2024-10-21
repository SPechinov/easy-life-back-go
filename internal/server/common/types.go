package common

import (
	"easy-life-back-go/internal/pkg/store_codes"
	"easy-life-back-go/pkg/store"
	"github.com/labstack/echo/v4"
)

type RoutesParams struct {
	Echo       *echo.Group
	Store      store.Store
	StoreCodes store_codes.StoreCodes
}
