package dumpling

import "regexp"

// Return a new Router instance
func New() *Router {
	r := &Router{}
	r.handlersMap = make(map[string]map[*regexp.Regexp]fn)
	r.chainedMiddlewares = nil
	return r
}

// Return a new HTTPContext instance
func newHTTPContext() *HTTPContext {
	c := &HTTPContext{}
	// use 200 by default
	c.respStatusCode = 200
	c.respHeaders = make(map[string]string)
	c.respContent = ""
	return c
}
