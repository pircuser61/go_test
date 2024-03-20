package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientOK(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}
	svr := httptest.NewServer(http.HandlerFunc(f))
	defer svr.Close()
	want := "Response:OK"
	got := request(svr.URL)
	if want != got {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func TestClientERR(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
	svr := httptest.NewServer(http.HandlerFunc(f))
	defer svr.Close()
	want := "HTTP 400"
	got := request(svr.URL)
	if want != got {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}
