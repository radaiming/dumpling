/*
Created © 2016-05-17 18:00 by @radaiming
*/

package dumpling

import (
	"fmt"
	"net/http"
)

type fn func() (int, map[string]string, string)

type Router struct {
	// many middlewares are not func(http.Handler) http.Handler
	// so let use chains them
	handlersMap        map[string]map[string]fn
	chainedMiddlewares http.Handler
}

func (r *Router) addRoute(path string, method string, f fn) {
	if _, ok := r.handlersMap[path]; !ok {
		r.handlersMap[path] = map[string]fn{method: f}
	} else {
		r.handlersMap[path][method] = f
	}
}

func (r *Router) Get(path string, f fn) {
	r.addRoute(path, "GET", f)
}

func (r *Router) Post(path string, f fn) {
	r.addRoute(path, "POST", f)
}

func (r *Router) Plug(h http.Handler) {
	r.chainedMiddlewares = h
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	f, ok := r.handlersMap[req.URL.Path][req.Method]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		statusCode, header, content := f()
		if header != nil {
			for k, v := range header {
				writer.Header().Set(k, v)
			}
		}
		writer.WriteHeader(statusCode)
		writer.Write([]byte(content))
	}
}

func (r *Router) Serve(addr string) {
	fmt.Println("now serving on " + addr)
	var final http.Handler
	if r.chainedMiddlewares != nil {
		final = r.chainedMiddlewares
	} else {
		final = r
	}
	// will call final's ServeHTTP
	err := http.ListenAndServe(addr, final)
	if err != nil {
		panic(err)
	}
}
