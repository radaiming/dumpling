/*
Created © 2016-05-17 22:32 by @radaiming
*/

package dumpling

import (
	"net/http"
)

type HTTPContext struct {
	respStatusCode int
	respHeaders    map[string]string
	respContent    string

	reqHeaders http.Header
}

func (h *HTTPContext) SetStatusCode(code int) {
	h.respStatusCode = code
}

func (h *HTTPContext) AddHeader(k string, v string) {
	h.respHeaders[k] = v
}

func (h *HTTPContext) Response(content string) {
	h.respContent = content
}

func (h *HTTPContext) GetHeader(key string) string {
	return h.reqHeaders.Get(key)
}
