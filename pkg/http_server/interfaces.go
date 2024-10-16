package http_server

import "net/http"

type HttpServer struct {
	*HttpGroup
	router *http.ServeMux
	port   string
}

type Next func()

type Middleware func(http.ResponseWriter, *http.Request, Next)

type HttpHandle interface {
	Handle(pattern string, handler http.HandlerFunc, middlewares ...Middleware)
}

type route struct {
	method      string
	pattern     string
	handler     http.HandlerFunc
	middlewares []Middleware
}

type HttpGroup struct {
	prefix          string
	parentHttpGroup *HttpGroup
	httpGroups      map[string]*HttpGroup
	routes          map[string]*route
	middlewares     []Middleware
}
