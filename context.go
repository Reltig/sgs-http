package sgshttp

import (
	"context"
	"encoding/json"
	"net/http"
)

type Response struct {
	handler http.HandlerFunc
}

type RequestHandler func(*Context) Response

type Context struct {
	w       http.ResponseWriter
	req     *http.Request
	context context.Context
}

func (ctx *Context) wrap(f func()) Response {
	return Response{
		handler: func(w http.ResponseWriter, r *http.Request) { f() },
	}
}

func (ctx *Context) ResponseText(text string) Response {
	return ctx.wrap(func() {
		ctx.w.Write([]byte(text))
	})
}

func (ctx *Context) ResponseJSON(result any) Response {
	return ctx.wrap(func() {
		ctx.w.Header().Set("content-type", "application/json")
		response, _ := json.Marshal(result)
		// TODO: handle error
		ctx.w.Write(response)
	})
}

func (ctx *Context) ResponseEmpty() Response {
	return ctx.wrap(func() {
		ctx.w.WriteHeader(204)
		ctx.w.Write(nil)
	})
}
