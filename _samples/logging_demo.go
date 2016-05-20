package main

import (
	"github.com/radaiming/dumpling"
	"github.com/radaiming/dumpling/middlewares"
)

func hello(ctx *dumpling.HTTPContext) {
	ctx.Response("hello world")
}

/*
> go run logging_demo.go
now serving on 127.0.0.1:9988
2016/05/18 12:57:08 127.0.0.1:54232 curl/7.43.0 "POST / HTTP/1.1" 404 0
2016/05/18 12:57:14 127.0.0.1:54232 curl/7.43.0 "GET / HTTP/1.1" 200 11
*/
func main() {
	r := dumpling.New()
	chainedMiddleware := middlewares.Logging(r)
	r.Plug(chainedMiddleware)
	r.Get("/", hello)
	r.Serve("127.0.0.1:9988")
}
