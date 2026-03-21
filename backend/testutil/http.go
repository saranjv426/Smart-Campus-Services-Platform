package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MustMarshal marshals a value to JSON and panics on failure.
func MustMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

// PerformRequest executes an HTTP request against a gin/http handler.
func PerformRequest(handler http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return PerformRawRequest(handler, req)
}

// NewJSONRequest builds an HTTP request and sets the JSON content type when a body is present.
func NewJSONRequest(t *testing.T, method, path string, body []byte) *http.Request {
	t.Helper()

	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}

// PerformRawRequest executes a preconfigured HTTP request against a gin/http handler.
func PerformRawRequest(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}
