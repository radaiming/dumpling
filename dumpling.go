package dumpling

func New() *Router {
	r := &Router{}
	r.handlersMap = make(map[string]map[string]fn)
	return r
}
