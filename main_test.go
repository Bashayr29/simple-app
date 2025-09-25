package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	// Call the handler with our request and recorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := "Hi there new pipeline build from github add webhooks :)"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandlerWithDifferentMethods(t *testing.T) {
	tests := []struct {
		name   string
		method string
	}{
		{"GET Request", "GET"},
		{"POST Request", "POST"},
		{"PUT Request", "PUT"},
		{"DELETE Request", "DELETE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handler)

			handler.ServeHTTP(rr, req)

			// The handler should respond to all HTTP methods with the same message
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code for %s: got %v want %v",
					tt.method, status, http.StatusOK)
			}

			expected := "Hi there new pipeline build from github add webhooks :)"
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body for %s: got %v want %v",
					tt.method, rr.Body.String(), expected)
			}
		})
	}
}

func TestHandlerResponseHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr, req)

	// Check that the Content-Type header is set correctly
	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/plain; charset=utf-8" {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, "text/plain; charset=utf-8")
	}
}

func BenchmarkHandler(b *testing.B) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)
		handler.ServeHTTP(rr, req)
	}
}