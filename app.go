package sgshttp

import (
	"context"
	"net/http"
)

type App struct {
	Port   string
	router *Router
}

func CreateApp() *App {
	return &App{
		router: CreateRouter("/"),
	}
}

func (app *App) Get(path string, handler RequestHandler) {
	app.router.Get(path, handler)
}

func (app *App) Post(path string, handler RequestHandler) {
	app.router.Post(path, handler)
}

func (app *App) Patch(path string, handler RequestHandler) {
	app.router.Patch(path, handler)
}

func (app *App) Delete(path string, handler RequestHandler) {
	app.router.Delete(path, handler)
}

func (app *App) Mount(router *Router) {
	app.router.Mount(router)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqHandler, err := app.router.resolvePath(req.Method, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	} else {
		ctx := &Context{w, req, context.Background()}
		responseHandler := reqHandler(ctx).handler
		responseHandler.ServeHTTP(ctx.w, ctx.req)
	}
}

func (app *App) Listen(port string) {
	app.Port = port
	http.ListenAndServe(app.Port, app)
}
