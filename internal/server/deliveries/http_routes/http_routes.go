package http_routes

import (
	"easy-life-back-go/internal/server/deliveries/http_routes/routes/auth"
	"easy-life-back-go/pkg/http_server_old"
)

type HttpRoutes struct {
	httpServer *http_server_old.HttpServer
}

func New(httpServer *http_server_old.HttpServer) *HttpRoutes {
	routes := &HttpRoutes{
		httpServer: httpServer,
	}

	http_routes_auth.NewRouter(routes.httpServer)

	return routes
}
