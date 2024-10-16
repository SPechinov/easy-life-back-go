package http_server_old

import "net/http"

func (s *HttpServer) Get(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodGet, pattern, handler, middlewares...)
}

func (s *HttpServer) Head(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodHead, pattern, handler, middlewares...)
}

func (s *HttpServer) Post(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodPost, pattern, handler, middlewares...)
}

func (s *HttpServer) Put(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodPut, pattern, handler, middlewares...)
}

func (s *HttpServer) Patch(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodPatch, pattern, handler, middlewares...)
}

func (s *HttpServer) Delete(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodDelete, pattern, handler, middlewares...)
}

func (s *HttpServer) Connect(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodConnect, pattern, handler, middlewares...)
}

func (s *HttpServer) Options(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodOptions, pattern, handler, middlewares...)
}

func (s *HttpServer) Trace(
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	s.AddRoute(http.MethodTrace, pattern, handler, middlewares...)
}
