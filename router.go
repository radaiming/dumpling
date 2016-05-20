/*
Created © 2016-05-17 18:00 by @radaiming
*/

package dumpling

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type fn func(*HTTPContext)

// The Router contains user defined path/handler map,
// the user chained middlewares is also inside
type Router struct {
	// map[HTTP Method][Compiled Regexp] -> fn
	handlersMap map[string]map[*regexp.Regexp]fn
	// many middlewares are not func(http.Handler) http.Handler
	// so let use chains them
	chainedMiddlewares http.Handler
}

func (r *Router) addRoute(path string, method string, f fn) {
	regexpPtr := regexp.MustCompile(path)
	if _, ok := r.handlersMap[method]; !ok {
		r.handlersMap[method] = map[*regexp.Regexp]fn{regexpPtr: f}
	} else {
		r.handlersMap[method][regexpPtr] = f
	}
}

// Add a handler for GET request under the path
func (r *Router) Get(path string, f fn) {
	r.addRoute(path, "GET", f)
}

// Add a handler for POST request under the path
func (r *Router) Post(path string, f fn) {
	r.addRoute(path, "POST", f)
}

// Plug in a chained middleware, if run multiple times,
// only the last call takes effect
func (r *Router) Plug(h http.Handler) {
	r.chainedMiddlewares = h
}

func initContext(ctx *HTTPContext, req *http.Request) {
	ctx.reqHeaders = req.Header
	ctx.reqArgs, _ = url.ParseQuery(req.URL.RawQuery)
	if req.Method == "POST" {
		contentType := req.Header.Get("Content-Type")
		if contentType == "application/x-www-form-urlencoded" {
			req.ParseForm()
			ctx.postForm = req.PostForm
		} else if strings.Index(contentType, "multipart/form-data") >= 0 {
			ctx.multipartStreamReader, _ = req.MultipartReader()
		}
	}
}

// This function implements http.Handler interface
func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	regFnMap, ok := r.handlersMap[req.Method]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	var f fn
	for reg, _ := range regFnMap {
		if reg.FindString(req.URL.Path) == req.URL.Path {
			f = regFnMap[reg]
			break
		}
	}
	if f == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	ctx := newHTTPContext()
	initContext(ctx, req)
	f(ctx)
	for k, v := range ctx.respHeaders {
		writer.Header().Set(k, v)
	}
	writer.WriteHeader(ctx.respStatusCode)
	writer.Write([]byte(ctx.respContent))
}

// Start serving on given address
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
