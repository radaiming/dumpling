package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/radaiming/dumpling"
)

func simpleAppend(h http.Handler) http.Handler {
	// http://www.alexedwards.net/blog/making-and-using-middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Write([]byte("\nmiddleware appended\n"))
	})
}

func hello() (int, map[string]string, string) {
	return 200, nil, "hello world"
}

func main() {
	r := dumpling.New()
	chainedMiddleware := handlers.CompressHandler(
		handlers.LoggingHandler(os.Stderr, simpleAppend(r)))
	r.Plug(chainedMiddleware)
	r.Get("/", hello)
	r.Serve("127.0.0.1:9988")
}
