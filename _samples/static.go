package main

import (
	"github.com/radaiming/dumpling"
	"github.com/radaiming/dumpling/middlewares"
)

func hello(ctx *dumpling.HTTPContext) {
	ctx.Response("hello world")
}

func main() {
	r := dumpling.New()
	chainedMiddleware := middlewares.ServeStatic("/static", "/tmp/", r)
	r.Plug(chainedMiddleware)
	r.Post("/static", hello)
	r.Serve("127.0.0.1:9988")
}
