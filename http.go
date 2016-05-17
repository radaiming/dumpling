/*
Created © 2016-05-17 22:32 by @radaiming
*/

package dumpling

type HTTPContext struct {
	respStatusCode int
	respHeaders    map[string]string
	respContent    string

	reqHeaders map[string]string
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
