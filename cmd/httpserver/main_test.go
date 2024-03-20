package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rw := httptest.NewRecorder()
	want := "OK"

	handleRoot(rw, req)

	resp := rw.Result()
	defer resp.Body.Close()
	var got string
	fmt.Fscan(resp.Body, &got)
	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}
