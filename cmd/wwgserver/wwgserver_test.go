package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HelloWorldHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	got, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	want := "Hello worlde!\n"
	if fmt.Sprintf("%s", got) != want {
		t.Fatalf("Expecting %q, got %q", want, got)
	}
}
