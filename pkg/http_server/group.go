package http_server

import (
	"log/slog"
	"net/http"
)

func CreateGroup(prefix string, parentHttpGroup *HttpGroup) *HttpGroup {
	return &HttpGroup{
		prefix:          prefix,
		parentHttpGroup: parentHttpGroup,
		httpGroups:      map[string]*HttpGroup{},
		routes:          map[string]*route{},
		middlewares:     []Middleware{},
	}
}

func (g *HttpGroup) Use(middlewares ...Middleware) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *HttpGroup) NewGroup(prefix string) *HttpGroup {
	httpGroup := CreateGroup(prefix, g)
	g.httpGroups[prefix] = httpGroup
	return httpGroup
}

func (g *HttpGroup) AddRoute(
	method string,
	pattern string,
	handler http.HandlerFunc,
	middlewares ...Middleware,
) {
	key := method + " " + pattern

	if _, existRoute := g.routes[key]; existRoute {
		slog.Error("Route already registered", "method", method, "route", g.getParentPrefixes(), "pattern", pattern)
		return
	}

	g.routes[key] = &route{
		handler:     handler,
		middlewares: middlewares,
		method:      method,
		pattern:     pattern,
	}
}

func (g *HttpGroup) getParentPrefixes() string {
	result := ""

	var getPrefix func(group *HttpGroup)

	getPrefix = func(group *HttpGroup) {
		if group == nil {
			return
		}

		result = group.prefix + result

		if group.parentHttpGroup != nil {
			getPrefix(group.parentHttpGroup)
		}
	}

	getPrefix(g)

	return result
}
