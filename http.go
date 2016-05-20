/*
Created © 2016-05-17 22:32 by @radaiming
*/

package dumpling

import (
	"mime/multipart"
	"net/http"
	"net/url"
)

// The HTTPContext is the only paramemter passed to
// user defined request handler. The HTTPContext
// represents both client request and server response,
// you could get content from or set content to it
type HTTPContext struct {
	// I should export these variables, following
	// function is too tedious
	respStatusCode int
	respHeaders    map[string]string
	respContent    string

	reqHeaders            http.Header
	reqArgs               url.Values
	postForm              url.Values
	multipartStreamReader *multipart.Reader
}

// Set response status code
func (h *HTTPContext) SetStatusCode(code int) {
	h.respStatusCode = code
}

// Add HTTP header for response
func (h *HTTPContext) AddHeader(k string, v string) {
	h.respHeaders[k] = v
}

// Set response content, for example, html string
func (h *HTTPContext) Response(content string) {
	h.respContent = content
}

// Get HTTP header of client request
func (h *HTTPContext) GetHeader(key string) string {
	return h.reqHeaders.Get(key)
}

// Get all request arguments from a GET request
func (h *HTTPContext) GetReqArgs() url.Values {
	return h.reqArgs
}

// Get request argument from a GET request
func (h *HTTPContext) GetReqArg(key string) string {
	return h.reqArgs.Get(key)
}

// Get all POST forms from an application/x-www-form-urlencoded request
func (h *HTTPContext) GetPostForms() url.Values {
	return h.postForm
}

// Get POST form value from an application/x-www-form-urlencoded request
func (h *HTTPContext) GetPostForm(key string) string {
	return h.postForm.Get(key)
}

// Get *multipart.Reader from multipart/form-data request
func (h *HTTPContext) GetMultipartStreamReader() *multipart.Reader {
	return h.multipartStreamReader
}
