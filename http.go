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
	// HTTP status code of server response
	RespStatusCode int

	// HTTP headers of server response
	RespHeaders map[string]string

	// HTTP content of server response
	RespContent string

	// HTTP header of client request
	ReqHeaders http.Header

	// HTTP request arguments of client GET request
	ReqArgs url.Values

	// HTTP POST forms of client application/x-www-form-urlencoded request
	PostForm url.Values

	// *multipart.Reader of client multipart/form-data request
	MultipartStreamReader *multipart.Reader
}

// Set response status code
func (h *HTTPContext) SetStatusCode(code int) {
	h.RespStatusCode = code
}

// Add HTTP header for response
func (h *HTTPContext) AddHeader(k string, v string) {
	h.RespHeaders[k] = v
}

// Set response content, for example, html string
func (h *HTTPContext) Response(content string) {
	h.RespContent = content
}

// Get HTTP header of client request
func (h *HTTPContext) GetHeader(key string) string {
	return h.ReqHeaders.Get(key)
}

// Get all request arguments from a GET request
func (h *HTTPContext) GetReqArgs() url.Values {
	return h.ReqArgs
}

// Get request argument from a GET request
func (h *HTTPContext) GetReqArg(key string) string {
	return h.ReqArgs.Get(key)
}

// Get all POST forms from an application/x-www-form-urlencoded request
func (h *HTTPContext) GetPostForms() url.Values {
	return h.PostForm
}

// Get POST form value from an application/x-www-form-urlencoded request
func (h *HTTPContext) GetPostForm(key string) string {
	return h.PostForm.Get(key)
}

// Get *multipart.Reader from multipart/form-data request
func (h *HTTPContext) GetMultipartStreamReader() *multipart.Reader {
	return h.MultipartStreamReader
}
