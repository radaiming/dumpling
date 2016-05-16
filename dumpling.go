package dumpling

import (
	"fmt"
	"net/http"
)

// return status code, header and content
type fn func() (int, map[string]string, string)

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

func Serve(addr string) {
	http.HandleFunc("/", handler)
	fmt.Println("now serving on " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
