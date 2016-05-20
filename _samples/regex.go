package main

import (
	"github.com/radaiming/dumpling"
)

func blabla(ctx *dumpling.HTTPContext) {
	ctx.Response("URL matches!")
}

func main() {
	r := dumpling.New()
	r.Get("/.+?/blabla/*", blabla)
	r.Serve("127.0.0.1:9988")
}
