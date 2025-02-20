package sgshttp

import (
	"net/http"
	"fmt"
)

type route struct {
	method string
	path string
	handler func(w http.ResponseWriter, req *http.Request)
}

type App struct {
	Routes []route
	Port int
}

func CreateApp() App {
	return App{}
}

func (app *App) initApp() {

}

func (app *App) addRoute(method string, path string, handler func(w http.ResponseWriter, req *http.Request)) {
	app.Routes = append(app.Routes, route{method, path, handler})
}

func (app *App) Get(path string, handler func(w http.ResponseWriter, req *http.Request)) {
	app.addRoute("GET", path, handler)
}

func (app *App) Post(path string, handler func(w http.ResponseWriter, req *http.Request)) {
	app.addRoute("POST", path, handler)
}

func (app *App) Listen(port int) {
	app.Port = port
	for _, route := range app.Routes {
		http.HandleFunc(route.path, route.handler)
	}
	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), nil)
}