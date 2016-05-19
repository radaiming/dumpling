package main

import (
	"github.com/radaiming/dumpling"
	"github.com/radaiming/dumpling/middlewares"
)

func main() {
	r := dumpling.New()
	chainedMiddleware := middlewares.ServeStatic("/static", "/tmp/")
	r.Plug(chainedMiddleware)
	r.Serve("127.0.0.1:9988")
}
