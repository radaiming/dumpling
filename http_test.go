package dumpling

import (
	"net/http"
	"net/url"
	"testing"
)

func TestSetStatusCode(t *testing.T) {
	c := newHTTPContext()
	c.SetStatusCode(500)
	if c.RespStatusCode != 500 {
		t.Error("Failed to set resopnse status code")
	}
}

func TestAddHeader(t *testing.T) {
	c := newHTTPContext()
	c.AddHeader("Content-Type", "plain/text")
	if c.RespHeaders["Content-Type"] != "plain/text" {
		t.Error("Failed to add response HTTP Header")
	}
}

func TestGetHeader(t *testing.T) {
	c := newHTTPContext()
	headers := http.Header{}
	headers.Add("Content-Type", "application/xml")
	c.ReqHeaders = headers
	if c.GetHeader("Content-Type") != "application/xml" {
		t.Error("Failed to get request HTTP header")
	}
}

func TestResponse(t *testing.T) {
	c := newHTTPContext()
	c.Response("blabla")
	if c.RespContent != "blabla" {
		t.Error("Failed to set response content")
	}
}

func TestGetReqArgs(t *testing.T) {
	c := newHTTPContext()
	args := url.Values{}
	args.Add("a", "b")
	args.Add("c", "d")
	c.ReqArgs = args
	if len(c.GetReqArgs()) != len(args) || c.GetReqArgs()["a"][0] != "b" || c.GetReqArgs()["c"][0] != "d" {
		t.Error("GetReqArgs() returns different values")
	}
	if c.GetReqArg("a") != "b" || c.GetReqArg("c") != "d" {
		t.Error("GetReqArg() returns different value")
	}
}
