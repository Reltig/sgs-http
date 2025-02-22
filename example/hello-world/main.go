package main

import (
	sgs "github.com/Reltig/sgs-http"
	"net/http"
)

func main() {
	app := sgs.CreateApp()

	app.Get("/hello", hello)

	app.Listen(":8000")
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}