package dumpling

func New() *Router {
	r := &Router{}
	r.handlersMap = make(map[string]map[string]fn)
	r.chainedMiddlewares = nil
	return r
}

func newHTTPContext() *HTTPContext {
	c := &HTTPContext{}
	// use 200 by default
	c.respStatusCode = 200
	c.respHeaders = make(map[string]string)
	c.respContent = ""
	return c
}
