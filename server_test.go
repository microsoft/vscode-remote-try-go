package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal("Failed to create 'GET /' request")
	}

	res := httptest.NewRecorder()

	want := "Hello, "

	handle(res, req)
	if got := res.Body.String(); want != got {
		t.Errorf("Handle() = %q, want = %q", got, want)
	}
}
