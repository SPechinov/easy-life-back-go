package http_routes_auth

import (
	"easy-life-back-go/pkg/http_server_old"
)

type HttpAuthRoutes struct {
	httpServer *http_server_old.HttpServer
	controller *Controller
}

func NewRouter(httpServer *http_server_old.HttpServer) *HttpAuthRoutes {
	httpRoutesAuth := &HttpAuthRoutes{
		httpServer: httpServer,
		controller: newController(),
	}

	httpRoutesAuth.startRouting()
	return httpRoutesAuth
}

func (routes *HttpAuthRoutes) startRouting() {
	routes.httpServer.Post("/auth/login", routes.controller.login)
	routes.httpServer.Post("/auth/registration", routes.controller.registration)
	routes.httpServer.Post("/auth/registration-success", routes.controller.registrationSuccess)
	routes.httpServer.Post("/auth/forgot-password", routes.controller.forgotPassword)
	routes.httpServer.Post("/auth/forgot-password-success", routes.controller.forgotPasswordSuccess)
}
