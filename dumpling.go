package dumpling

import "regexp"

// Return a new Router instance
func New() *Router {
	r := &Router{}
	r.HandlersMap = make(map[string]map[*regexp.Regexp]fn)
	r.ChainedMiddlewares = nil
	return r
}

// Return a new HTTPContext instance
func newHTTPContext() *HTTPContext {
	c := &HTTPContext{}
	// use 200 by default
	c.RespStatusCode = 200
	c.RespHeaders = make(map[string]string)
	c.RespContent = ""
	return c
}
