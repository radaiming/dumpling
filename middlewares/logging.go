package middlewares

import (
	"log"
	"net/http"
	"net/http/httptest"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// https://justinas.org/writing-http-middleware-in-go/
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
		log.Printf("%s %s \"%s %s %s\" %d %d\n",
			r.RemoteAddr, r.UserAgent(), r.Method, r.URL.Path,
			r.Proto, rec.Code, rec.Body.Len())
	})
}
