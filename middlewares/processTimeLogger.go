package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Middleware for logging process time of client requests,
// log will be printed to stdout of server side
func ProcessTimeLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startPoint := time.Now().Nanosecond()
		h.ServeHTTP(w, r)
		elapsed := float64(time.Now().Nanosecond()-startPoint) / 1000000.0
		log.Printf("%s \"%s %s %s\" %.3f ms\n",
			r.RemoteAddr, r.Method, r.URL.Path,
			r.Proto, elapsed)
	})
}
