package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForward(t *testing.T) {
	// Create a GET fake request to pass to the handler
	req, err := http.NewRequest("GET", "/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response and instantiate the handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Forward)

	// Pass in the Request and ResponseRecorder to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "/path"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
