package main

import (
	sgs "github.com/Reltig/sgs-http"
)

func main() {
	app := sgs.CreateApp()

	app.Get("/hello", hello)
	app.Post("/hello/world", helloWorld)

	app.Listen(":8000")

}

func hello(ctx *sgs.Context) sgs.Response {
	return ctx.ResponseText("Hello")
}

func helloWorld(ctx *sgs.Context) sgs.Response {
	return ctx.ResponseText("Hello world")
}
