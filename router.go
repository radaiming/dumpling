/*
Created © 2016-05-17 18:00 by @radaiming
*/

package dumpling

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

type fn func(*HTTPContext)

// The Router contains user defined path/handler map,
// the user chained middlewares is also inside
type Router struct {
	// A map for user defined handler and path:
	// map[HTTP Method][Compiled Regexp] -> fn
	HandlersMap map[string]map[*regexp.Regexp]fn

	// many middlewares are not func(http.Handler) http.Handler,
	// so please chain them manually, then pass in
	ChainedMiddlewares http.Handler
}

func (r *Router) addRoute(path string, method string, f fn) {
	regexpPtr := regexp.MustCompile(path)
	if _, ok := r.HandlersMap[method]; !ok {
		r.HandlersMap[method] = map[*regexp.Regexp]fn{regexpPtr: f}
	} else {
		r.HandlersMap[method][regexpPtr] = f
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
	r.ChainedMiddlewares = h
}

func initContext(ctx *HTTPContext, req *http.Request) {
	ctx.ReqHeaders = req.Header
	ctx.ReqArgs, _ = url.ParseQuery(req.URL.RawQuery)
	if req.Method == "POST" {
		contentType := req.Header.Get("Content-Type")
		if contentType == "application/x-www-form-urlencoded" {
			req.ParseForm()
			ctx.PostForm = req.PostForm
		} else {
			err := req.ParseMultipartForm(32 << 20)
			if err == nil {
				ctx.MultipartForm = req.MultipartForm
			}
		}
	}
}

// This function implements http.Handler interface
func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	regFnMap, ok := r.HandlersMap[req.Method]
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
	for k, v := range ctx.RespHeaders {
		writer.Header().Set(k, v)
	}
	writer.WriteHeader(ctx.RespStatusCode)
	writer.Write([]byte(ctx.RespContent))
}

// Start serving on given address
func (r *Router) Serve(addr string) {
	var final http.Handler
	if r.ChainedMiddlewares != nil {
		final = r.ChainedMiddlewares
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
