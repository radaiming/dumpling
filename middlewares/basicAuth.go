package middlewares

import "net/http"

// Middleware for basic authentication on client requests
func BasicAuth(username string, password string, h http.Handler) http.Handler {
	// https://golang.org/src/net/http/request.go?s=20802:20868#L633
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUsername, reqPassword, ok := r.BasicAuth()
		if !ok || reqUsername != username || reqPassword != password {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
