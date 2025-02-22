package sgshttp

import (
	"context"
	"net/http"
)

type Response struct {
	handler http.HandlerFunc
}

type RequestHandler func(*RequestContext) Response

type RequestContext struct {
	w   http.ResponseWriter
	req *http.Request
	context context.Context
}

func (ctx *RequestContext) wrap(f func()) Response {
	return Response{
		handler: func(w http.ResponseWriter, r *http.Request) { f() },
	}
}

func (ctx *RequestContext) ResponseText(text string) Response {
	return ctx.wrap(func() {
		ctx.w.Header().Set("content-type", "text/plain; charset=utf-8")
		ctx.w.Write([]byte(text))
	})
}
