package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func performMiddlewareRequest(t *testing.T, handler http.Handler, method, path, headerName, headerValue string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(method, path, nil)
	if headerName != "" {
		req.Header.Set(headerName, headerValue)
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}

func init() {
	gin.SetMode(gin.TestMode)
}
