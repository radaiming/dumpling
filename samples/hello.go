/*
Created © 2016-05-16 19:31 by @radaiming
*/

package main

import (
	"github.com/radaiming/dumpling"
)

func hello() (int, map[string]string, string) {
	return 200, nil, "hello world"
}

func main() {
	r := dumpling.New()
	r.Get("/", hello)
	r.Post("/", hello)
	r.Serve("127.0.0.1:9988")
}
