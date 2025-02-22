package sgshttp

import (
	"errors"
	"net/http"
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
	handler http.Handler
}

func (rt *Router) addRoute(method string, path string, handler http.Handler) {
	rt.routes = append(rt.routes, Route{method, path, handler})
}

func (rt *Router) resolvePath(path string) (http.Handler, error) {
	return nil, ErrRouteNotFound
}
