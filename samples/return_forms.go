/*
> curl -d "a=1&b=2&c=3" 'http://127.0.0.1:9988/'
b -> 2
c -> 3
a -> 1
*/
package main

import (
	"fmt"
	"strings"

	"github.com/radaiming/dumpling"
)

func readAndReturn(ctx *dumpling.HTTPContext) {
	var returnContent string
	for k, v := range ctx.GetPostForms() {
		returnContent += fmt.Sprintf("%s -> %s\n", k, strings.Join(v, " "))
	}
	ctx.Response(returnContent)
}

func main() {
	r := dumpling.New()
	r.Post("/", readAndReturn)
	r.Serve("127.0.0.1:9988")
}
