/*
Created © 2016-05-17 01:26 by @radaiming
*/

package main

import "github.com/radaiming/dumpling"

func redir(ctx *dumpling.HTTPContext) {
	ctx.SetStatusCode(302)
	ctx.AddHeader("Location", "https://google.com")
}

func main() {
	r := dumpling.New()
	r.Get("/", redir)
	r.Serve("127.0.0.1:9988")
}
