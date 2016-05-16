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
	dumpling.AddRoute("/", "GET", hello)
	dumpling.AddRoute("/", "POST", hello)
	dumpling.Serve("127.0.0.1:9988")
}
