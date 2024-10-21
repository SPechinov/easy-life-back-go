package auth

import (
	"easy-life-back-go/internal/server/common"
	"easy-life-back-go/internal/server/routes/auth/controller"
	"easy-life-back-go/internal/server/routes/auth/redis"
	"easy-life-back-go/internal/server/routes/auth/routes"
	"easy-life-back-go/internal/server/routes/auth/service"
)

func RegisterRoutes(params *common.RoutesParams) {
	r := redis.NewRedis(params.Store, params.StoreCodes)
	s := service.NewService(r)
	c := controller.NewController(s)
	routes.RegisterRoutes(params.Echo, c)
}
