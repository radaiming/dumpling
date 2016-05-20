package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	returnSecret := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	chainedMiddleware := BasicAuth("user001", "token001", returnSecret)

	unAuthedReq, err := http.NewRequest("GET", "http://127.0.0.1:9988/", nil)
	if err != nil {
		t.Error("Failed to construct request")
	}
	w := httptest.NewRecorder()
	chainedMiddleware.ServeHTTP(w, unAuthedReq)
	if w.Code != 401 {
		t.Error("Unauthorized request get non 401 response")
	}

	authedReq, err := http.NewRequest("GET", "http://127.0.0.1:9988/", nil)
	if err != nil {
		t.Error("Failed to construct request")
	}
	authedReq.Header.Add("Authorization", "Basic dXNlcjAwMTp0b2tlbjAwMQ==")
	w = httptest.NewRecorder()
	chainedMiddleware.ServeHTTP(w, authedReq)
	if w.Code != 200 {
		t.Error("Authorized request get non 200 response")
	}
}
