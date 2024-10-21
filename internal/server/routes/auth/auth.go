package auth

import (
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/internal/server/routes/auth/controller"
	"easy-life-back-go/internal/server/routes/auth/routes"
	"easy-life-back-go/internal/server/routes/auth/service"
	"easy-life-back-go/internal/server/routes/auth/store"
)

func RegisterRoutes(params *common.RoutesParams) {
	st := store.NewStore(params.Store, params.StoreCodes)
	s := service.NewService(st)
	c := controller.NewController(s)
	routes.RegisterRoutes(params.Echo, c)
}
