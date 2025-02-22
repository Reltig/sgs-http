package sgshttp

import (
	"net/http"
)

type App struct {
	Port string
	router Router
}

func CreateApp() App {
	return App{}
}

func (app *App) Get(path string, handler http.HandlerFunc) {
	app.router.addRoute("GET", path, handler)
}

func (app *App) Post(path string, handler http.HandlerFunc) {
	app.router.addRoute("POST", path, handler)
}

func (app *App) Patch(path string, handler http.HandlerFunc) {
	app.router.addRoute("PATCH", path, handler)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, err := app.router.resolvePath(req.Method, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	} else {
		handler.ServeHTTP(w, req)
	}
}

func (app *App) Listen(port string) {
	app.Port = port
	http.ListenAndServe(app.Port, app)
}
