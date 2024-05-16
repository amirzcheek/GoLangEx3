package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMaxRetries(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer mockServer.Close()

	retryCount = 3

	maxRetries = 3

	input := "test input"
	_, err := getChatResponse(input)

	expectedError := "max retry limit reached"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected '%s' error, got '%v'", expectedError, err)
	}
}
