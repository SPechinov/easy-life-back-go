package http_server_old

import (
	"net/http"
)

func New(port string) *HttpServer {
	mux := http.NewServeMux()

	return &HttpServer{
		router:      mux,
		port:        port,
		middlewares: []Middleware{},
		routes:      map[string]route{},
		headers:     map[string]string{},
	}
}
