package sgshttp

import (
	"errors"
)

var (
	ErrRouteNotFound = errors.New("Route not found")
)

type Router struct {
	routes []Route
}

type Route struct {
	method  string
	path    string
	handler RequestHandler
}

func (rt *Router) addRoute(method string, path string, handler RequestHandler) {
	rt.routes = append(rt.routes, Route{method, path, handler})
}

func (rt *Router) resolvePath(method string, path string) (RequestHandler, error) {
	for _, route := range rt.routes {
		if route.method == method && route.path == path {
			return route.handler, nil
		}
		// TODO: Method no allowed
	}
	return nil, ErrRouteNotFound
}
