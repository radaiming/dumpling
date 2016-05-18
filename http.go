/*
Created © 2016-05-17 22:32 by @radaiming
*/

package dumpling

import (
	"mime/multipart"
	"net/http"
	"net/url"
)

type HTTPContext struct {
	respStatusCode int
	respHeaders    map[string]string
	respContent    string

	reqHeaders            http.Header
	reqArgs               url.Values
	postForm              url.Values
	multipartStreamReader *multipart.Reader
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

func (h *HTTPContext) GetReqArgs() url.Values {
	return h.reqArgs
}

func (h *HTTPContext) GetReqArg(key string) string {
	return h.reqArgs.Get(key)
}

func (h *HTTPContext) GetPostForms() url.Values {
	return h.postForm
}

func (h *HTTPContext) GetPostForm(key string) string {
	return h.postForm.Get(key)
}

func (h *HTTPContext) GetMultipartStreamReader() *multipart.Reader {
	return h.multipartStreamReader
}
