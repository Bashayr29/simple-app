package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerWritesGreeting(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	if body := rec.Body.String(); body != greeting {
		t.Fatalf("expected body %q, got %q", greeting, body)
	}
}

func TestNewMuxRoutesRootToHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if body := rec.Body.String(); body != greeting {
		t.Fatalf("expected body %q, got %q", greeting, body)
	}
}

func TestNewMuxRoutesSubpathsToHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	newMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if body := rec.Body.String(); body != greeting {
		t.Fatalf("expected body %q, got %q", greeting, body)
	}
}

func TestRunUsesConfiguredAddressAndMux(t *testing.T) {
	originalListenAndServe := listenAndServe
	t.Cleanup(func() {
		listenAndServe = originalListenAndServe
	})

	called := false
	expectedErr := errors.New("server stopped")

	listenAndServe = func(addr string, handler http.Handler) error {
		called = true

		if addr != ":8080" {
			t.Fatalf("expected address %q, got %q", ":8080", addr)
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
		}

		if body := rec.Body.String(); body != greeting {
			t.Fatalf("expected body %q, got %q", greeting, body)
		}

		return expectedErr
	}

	err := run(":8080")
	if !called {
		t.Fatal("expected listenAndServe to be called")
	}

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}
}
