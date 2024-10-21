package auth

import (
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/internal/server/routes/auth/controller"
	"easy-life-back-go/internal/server/routes/auth/routes"
	"easy-life-back-go/internal/server/routes/auth/service"
	"easy-life-back-go/internal/server/routes/auth/store"
	"easy-life-back-go/internal/server/routes/auth/validation"
)

func RegisterRoutes(params *common.RoutesParams) {
	st := store.NewStore(params.Store, params.StoreCodes)
	s := service.NewService(st)
	v := validation.NewValidator()
	c := controller.NewController(s, v)
	routes.RegisterRoutes(params.Echo, c)
}
