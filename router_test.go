package dumpling

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"
)

func hello(ctx *HTTPContext) {
	ctx.Response("hello world")
}

func simpleAppend(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Write([]byte("\nmiddleware appended\n"))
	})
}

func TestAddPost(t *testing.T) {
	r := New()
	r.Post("/abc", hello)
	if len(r.HandlersMap["POST"]) != 1 {
		t.Error("Failed to add POST route")
	}
	for k, v := range r.HandlersMap["POST"] {
		if k.String() != "/abc" {
			t.Error("Compiled regexp returns a different source reg text")
		}
		ctx := newHTTPContext()
		v(ctx)
		if ctx.RespContent != "hello world" {
			t.Error("The function in added to HandlersMap is not returning what expected")
		}
	}
}

func TestAddGet(t *testing.T) {
	r := New()
	r.Get("/abc", hello)
	if len(r.HandlersMap["GET"]) != 1 {
		t.Error("Failed to add GET route")
	}
	for k, v := range r.HandlersMap["GET"] {
		if k.String() != "/abc" {
			t.Error("Compiled regexp returns a different source reg text")
		}
		ctx := newHTTPContext()
		v(ctx)
		if ctx.RespContent != "hello world" {
			t.Error("The function in added to HandlersMap is not returning what expected")
		}
	}
}

func TestPlug(t *testing.T) {
	r := New()
	r.Plug(simpleAppend(r))
	if _, ok := r.ChainedMiddlewares.(http.Handler); !ok {
		t.Error("Failed to add middleware")
	}
}

func TestInitContext(t *testing.T) {
	for _, method := range []string{"GET", "POST"} {
		ctx := newHTTPContext()
		req, err := http.NewRequest(method, "http://127.0.0.1:9988/abc?a=123", nil)
		if err != nil {
			t.Error("Failed to create Request")
		}
		req.Header.Add("User-Agent", "go testing")
		initContext(ctx, req)
		if ctx.GetHeader("User-Agent") != "go testing" {
			t.Error("HTTP Header mismatch")
		}
		if ctx.GetReqArg("a") != "123" {
			t.Error("HTTP query args mismatch")
		}
	}

	// test application/x-www-form-urlencoded
	ctx := newHTTPContext()
	req, err := http.NewRequest("POST", "http://127.0.0.1:9988/abc?a=123", nil)
	if err != nil {
		t.Error("Failed to create Request")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// PostForm is nil at initial
	req.PostForm = url.Values{}
	req.PostForm.Add("b", "234")
	initContext(ctx, req)
	if ctx.GetHeader("Content-Type") != "application/x-www-form-urlencoded" {
		t.Error("HTTP Header mismatch")
	}
	if ctx.GetReqArg("a") != "123" {
		t.Error("HTTP query args mismatch")
	}
	if ctx.GetPostForm("b") != "234" {
		t.Error("HTTP Post form mismatch")
	}

	// test multipart/form-data
	// https://matt.aimonetti.net/posts/2013/07/01/golang-multipart-file-upload-example/
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writingPart, err := writer.CreateFormField("some..field...name")
	if err != nil {
		t.Error("Failed to create form")
	}
	_, err = writingPart.Write([]byte("some...string"))
	if err != nil {
		t.Error("Error writing to multipart section")
	}
	if writer.Close() != nil {
		t.Error("Error closing multipart writer")
	}

	ctx = newHTTPContext()
	req, err = http.NewRequest("POST", "http://127.0.0.1:9988/", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	initContext(ctx, req)
	multipartStreamReader := ctx.GetMultipartStreamReader()
	readingPart, err := multipartStreamReader.NextPart()
	buffer := make([]byte, 1024)
	n, err := readingPart.Read(buffer)
	if err != nil {
		t.Error("Error reading from multipart section")
	}
	if string(buffer[0:n]) != "some...string" {
		t.Error("Get wrong string from multipart reader")
	}
}
