package http_server_old

import (
	"net/http"
)

func (s *HttpServer) startRouter() {
	for key, route := range s.routes {
		s.router.HandleFunc(key, func(w http.ResponseWriter, r *http.Request) {
			// Set default headers
			for header, value := range s.headers {
				w.Header().Set(header, value)
			}

			var next func()
			pointer := 0

			// Go through middlewares
			middlewares := append(s.middlewares, route.middlewares...)
			lenMiddlewares := len(middlewares)

			next = func() {
				if pointer < lenMiddlewares {
					pointer++
					middlewares[pointer-1](w, r, next)
				} else {
					route.handler(w, r)
				}
			}

			next()
		})
	}
}
