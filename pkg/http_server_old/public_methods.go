package http_server_old

import (
	"log/slog"
	"net"
	"net/http"
	"os"
)

func (s *HttpServer) Run() {
	slog.Info("Starting http server...")
	l, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		panic(err.Error())
		return
	}

	s.startRouter()

	slog.Info("🐱‍💻 Server started on", l.Addr().String())

	if err := http.Serve(l, s.router); err != nil {
		slog.Info("Server closed: %s\n", err)
	}
	os.Exit(1)
}

func (s *HttpServer) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *HttpServer) SetGlobalPrefix(prefix string) {
	s.globalPrefix = prefix
}

func (s *HttpServer) SetGlobalHeader(header, value string) {
	s.headers[header] = value
}

func (s *HttpServer) AddRoute(
	method string,
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	key := method + " " + s.globalPrefix + pattern

	if _, existRoute := s.routes[key]; existRoute {
		slog.Error("Route already registered", "method", method, "pattern", pattern, "key", key)
		return
	}

	s.routes[key] = route{
		handler:     handler,
		middlewares: middlewares,
	}
}
