package main

import (
	"math/rand"
	"time"

	"github.com/radaiming/dumpling"
	"github.com/radaiming/dumpling/middlewares"
)

func hello(ctx *dumpling.HTTPContext) {
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	ctx.Response("hello world")
}

/*
> go run process_time_logger_demo.go
now serving on 127.0.0.1:9988
2016/05/19 01:50:53 127.0.0.1:49755 "POST / HTTP/1.1" 1.270 ms
*/
func main() {
	r := dumpling.New()
	r.Plug(middlewares.ProcessTimeLogger(r))
	r.Get("/", hello)
	r.Serve("127.0.0.1:9988")
}
