package main

import (
	sgs "github.com/Reltig/sgs-http"
)

func main() {
	app := sgs.CreateApp()
	router := sgs.CreateRouter("/mounted")
	router.Get("/hello", func(ctx *sgs.Context) sgs.Response {
		return ctx.ResponseText("hello")
	})
	app.Mount(router)
	app.Listen(":8000")
}