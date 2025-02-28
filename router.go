package sgshttp

import (
	"errors"
)

var (
	ErrRouteNotFound = errors.New("Route not found")
)

type Router struct {
	BasePath string
	routes []Route
}

type Route struct {
	method  string
	path    string
	handler RequestHandler
}

func CreateRouter(BasePath string) *Router {
	return &Router{
		BasePath,
		[]Route{},
	}
}

func (router *Router) addRoute(method string, path string, handler RequestHandler) {
	router.routes = append(router.routes, Route{method, path, handler})
}

func (router *Router) Get(path string, handler RequestHandler) {
	router.addRoute("GET", path, handler)
}

func (router *Router) Post(path string, handler RequestHandler) {
	router.addRoute("POST", path, handler)
}

func (router *Router) Patch(path string, handler RequestHandler) {
	router.addRoute("PATCH", path, handler)
}

func (router *Router) Delete(path string, handler RequestHandler) {
	router.addRoute("DELETE", path, handler)
}

func (router *Router) Mount(mountedRouter *Router) {
	for _, route := range mountedRouter.routes {
		router.addRoute(route.method, mountedRouter.BasePath + route.path, route.handler)
	}
}

func (router *Router) resolvePath(method string, path string) (RequestHandler, error) {
	for _, route := range router.routes {
		if route.method == method && route.path == path {
			return route.handler, nil
		}
		// TODO: Method no allowed
	}
	return nil, ErrRouteNotFound
}
