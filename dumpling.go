package dumpling

func New() *Router {
	r := &Router{}
	r.handlersMap = make(map[string]map[string]fn)
	r.chainedMiddlewares = nil
	return r
}
