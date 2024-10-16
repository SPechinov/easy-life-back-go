package http_server_old

import "net/http"

type Next func()

type Middleware func(http.ResponseWriter, *http.Request, Next)

type HttpHandle interface {
	Handle(pattern string, handler http.HandlerFunc, middlewares ...Middleware)
}

type HttpServer struct {
	router       *http.ServeMux
	port         string
	globalPrefix string
	routes       map[string]route
	headers      map[string]string
	middlewares  []Middleware
}

type route struct {
	handler     http.HandlerFunc
	middlewares []Middleware
}
