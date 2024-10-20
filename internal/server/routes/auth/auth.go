package auth

import (
	"easy-life-back-go/internal/server"
	"easy-life-back-go/internal/server/routes/auth/controller"
	"easy-life-back-go/internal/server/routes/auth/redis"
	"easy-life-back-go/internal/server/routes/auth/routes"
	"easy-life-back-go/internal/server/routes/auth/service"
)

func RegisterRoutes(params *server.RoutesParams) {
	r := redis.NewRedis(params.redis)
	s := service.NewService(r)
	c := controller.NewController(s)
	routes.RegisterRoutes(params.echo, c)
}
