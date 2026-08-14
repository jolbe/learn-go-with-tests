package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetData(t *testing.T) {
	t.Run("GOOD EXAMPLE: when you don't get a 200 you get a status error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer srv.Close()

		_, err := DumbGetter(srv.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		// 1. option via type assertion
		got, isStatusErr := err.(BadStatusError)
		if !isStatusErr {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		// 2. option via errors.As
		var wantErr BadStatusError
		if !errors.As(err, &wantErr) {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		want := BadStatusError{URL: srv.URL, Status: http.StatusTeapot}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("BAD EXAMPLE: when you don't get a 200 you get a status error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer srv.Close()

		_, err := DumbGetter(srv.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		want := fmt.Sprintf("did not get 200 from %s, got %d", srv.URL, http.StatusTeapot)
		got := err.Error()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
