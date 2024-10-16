package http_server

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

func New(port, prefix string) *HttpServer {
	router := http.NewServeMux()

	return &HttpServer{
		HttpGroup: CreateGroup(prefix, nil),
		router:    router,
		port:      port,
	}
}

func (s *HttpServer) Run() error {
	s.startRouting()
	slog.Info("Starting http server...")
	l, err := net.Listen("tcp", ":"+s.port)

	if err != nil {
		return err
	}

	slog.Info("🐱‍💻 Server started on", l.Addr().String())

	if err := http.Serve(l, s.router); err != nil {
		return err
	}

	return nil
}

func (s *HttpServer) startRouting() {
	var start func(group *HttpGroup, prefix string, middlewares []Middleware)

	start = func(group *HttpGroup, prefix string, middlewares []Middleware) {
		for _, route := range group.routes {
			// Create correct pattern
			pattern := route.method + " " + prefix + route.pattern
			fmt.Println(pattern)

			// Wrap in middlewares
			middlewares := append(middlewares, route.middlewares...)
			lenMiddlewares := len(middlewares)

			handler := func(w http.ResponseWriter, r *http.Request) {
				var next func()
				pointer := 0

				next = func() {
					if pointer < lenMiddlewares {
						pointer++
						middlewares[pointer-1](w, r, next)
					} else {
						route.handler(w, r)
					}
				}

				next()
			}

			s.router.HandleFunc(pattern, handler)
		}

		if group.httpGroups != nil {
			for _, group := range group.httpGroups {
				start(
					group,
					prefix+group.prefix,
					append(middlewares, group.middlewares...),
				)
			}
		}
	}

	start(s.HttpGroup, s.HttpGroup.prefix, s.HttpGroup.middlewares)
}
