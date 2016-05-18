/*
Created © 2016-05-17 18:00 by @radaiming
*/

package dumpling

import (
	"fmt"
	"net/http"
	"net/url"
)

type fn func(*HTTPContext)

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

func initContext(ctx *HTTPContext, req *http.Request) {
	ctx.reqHeaders = req.Header
	ctx.reqArgs, _ = url.ParseQuery(req.URL.RawQuery)
	if req.Method == "POST" {
		if req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			req.ParseForm()
			ctx.postForm = req.PostForm
		}
	}
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	f, ok := r.handlersMap[req.URL.Path][req.Method]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		ctx := newHTTPContext()
		initContext(ctx, req)
		f(ctx)
		for k, v := range ctx.respHeaders {
			writer.Header().Set(k, v)
		}
		writer.WriteHeader(ctx.respStatusCode)
		writer.Write([]byte(ctx.respContent))
	}
}

func (r *Router) Serve(addr string) {
	var final http.Handler
	if r.chainedMiddlewares != nil {
		final = r.chainedMiddlewares
	} else {
		final = r
	}
	// will call final's ServeHTTP
	fmt.Println("now serving on " + addr)
	err := http.ListenAndServe(addr, final)
	if err != nil {
		panic(err)
	}
}
