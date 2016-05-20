package main

import (
	"github.com/radaiming/dumpling"
	"github.com/radaiming/dumpling/middlewares"
)

func returnSecret(ctx *dumpling.HTTPContext) {
	ctx.Response("Top Secret Content!")
}

func main() {
	r := dumpling.New()
	// run curl -H 'Authorization: Basic dXNlcjAwMTp0b2tlbjAwMQ==' to pass auth
	chainedMiddleware := middlewares.BasicAuth("user001", "token001", r)
	r.Plug(chainedMiddleware)
	r.Get("/", returnSecret)
	r.Serve("127.0.0.1:9988")
}
