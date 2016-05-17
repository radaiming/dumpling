/*
Created © 2016-05-17 01:26 by @radaiming
*/

package main

import "github.com/radaiming/dumpling"

func redir() (int, map[string]string, string) {
	headers := map[string]string{"Location": "https://google.com"}
	return 301, headers, ""
}

func main() {
	r := dumpling.New()
	r.Get("/", redir)
	r.Serve("127.0.0.1:9988")
}
