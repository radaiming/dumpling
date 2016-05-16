package dumpling

import (
	"fmt"
	"net/http"
)

type fn func() string

var handlersMap = make(map[string]map[string]fn)

func AddRoute(route string, method string, handlerFunc fn) {
	if _, ok := handlersMap[route]; !ok {
		handlersMap[route] = map[string]fn{method: handlerFunc}
	} else {
		handlersMap[route][method] = handlerFunc
	}
}

func handler(writer http.ResponseWriter, req *http.Request) {
	f, ok := handlersMap[req.URL.Path][req.Method]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.Write([]byte(f()))
	}
}

func Serve(addr string) {
	http.HandleFunc("/", handler)
	fmt.Println("now serving on " + addr)
	http.ListenAndServe(addr, nil)
}
