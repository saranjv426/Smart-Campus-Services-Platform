package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDatabasePathUsesEnvOverride(t *testing.T) {
	t.Setenv("DB_PATH", "custom/test.db")

	if got := databasePath(); got != "custom/test.db" {
		t.Fatalf("expected custom/test.db, got %s", got)
	}
}

func TestDatabasePathFallsBackToDefault(t *testing.T) {
	if err := os.Unsetenv("DB_PATH"); err != nil {
		t.Fatalf("failed to unset DB_PATH: %v", err)
	}

	if got := databasePath(); got != "data/smart_campus.db" {
		t.Fatalf("expected data/smart_campus.db, got %s", got)
	}
}

func TestCorsMiddlewareHandlesOptionsRequests(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(corsMiddleware())
	router.OPTIONS("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, err := http.NewRequest(http.MethodOptions, "/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rec.Code)
	}

	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatalf("expected Access-Control-Allow-Origin header to be set, got %q", rec.Header().Get("Access-Control-Allow-Origin"))
	}
}
