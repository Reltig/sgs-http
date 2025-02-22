package sgshttp

import (
	"context"
	"net/http"
)

type App struct {
	Port   string
	router Router
}

func CreateApp() *App {
	return &App{}
}

func (app *App) Get(path string, handler RequestHandler) {
	app.router.addRoute("GET", path, handler)
}

func (app *App) Post(path string, handler RequestHandler) {
	app.router.addRoute("POST", path, handler)
}

func (app *App) Patch(path string, handler RequestHandler) {
	app.router.addRoute("PATCH", path, handler)
}

func (app *App) Delete(path string, handler RequestHandler) {
	app.router.addRoute("DELETE", path, handler)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqHandler, err := app.router.resolvePath(req.Method, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	} else {
		ctx := &RequestContext{w, req, context.Background()}
		responseHandler := reqHandler(ctx).handler
		responseHandler.ServeHTTP(ctx.w, ctx.req)
	}
}

func (app *App) Listen(port string) {
	app.Port = port
	http.ListenAndServe(app.Port, app)
}
