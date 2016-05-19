package dumpling

import "testing"

func TestNew(t *testing.T) {
	r := New()
	if r == nil {
		t.Error("dumpling.New() should not return nil")
	}
}

func TestNewHTTPContext(t *testing.T) {
	c := newHTTPContext()
	if c == nil {
		t.Error("newHTTPContext() should not return nil")
	}
}
