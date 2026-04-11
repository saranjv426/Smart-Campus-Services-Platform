package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
)

func setupRouterTest(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	r := gin.New()
	RegisterAPIRoutes(r, db)
	return r
}

func TestProtectedRoutesRequireAuth(t *testing.T) {
	r := setupRouterTest(t)

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{name: "auth logout", method: http.MethodPost, path: "/api/auth/logout"},
		{name: "auth refresh", method: http.MethodPost, path: "/api/auth/refresh"},
		{name: "user update", method: http.MethodPut, path: "/api/users/123"},
		{name: "service create", method: http.MethodPost, path: "/api/services"},
		{name: "service update", method: http.MethodPut, path: "/api/services/123"},
		{name: "service toggle active", method: http.MethodPatch, path: "/api/services/123/active"},
		{name: "booking create", method: http.MethodPost, path: "/api/bookings"},
		{name: "booking update", method: http.MethodPut, path: "/api/bookings/123"},
		{name: "booking cancel status", method: http.MethodPatch, path: "/api/bookings/123/status"},
		{name: "approval approve", method: http.MethodPut, path: "/api/approval/bookings/123/approve"},
		{name: "notification read", method: http.MethodPut, path: "/api/notifications/123/read"},
		{name: "review create", method: http.MethodPost, path: "/api/reviews"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := testutil.PerformRequest(r, tt.method, tt.path, nil)
			if resp.Code != http.StatusUnauthorized {
				t.Fatalf("expected status %d, got %d body=%s", http.StatusUnauthorized, resp.Code, resp.Body.String())
			}
		})
	}
}

func TestPublicRoutesRemainAccessibleWithoutAuth(t *testing.T) {
	r := setupRouterTest(t)

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{name: "register", method: http.MethodPost, path: "/api/auth/register"},
		{name: "login", method: http.MethodPost, path: "/api/auth/login"},
		{name: "service list", method: http.MethodGet, path: "/api/services"},
		{name: "service get", method: http.MethodGet, path: "/api/services/123"},
		{name: "service category", method: http.MethodGet, path: "/api/services/category/library"},
		{name: "booking get", method: http.MethodGet, path: "/api/bookings/123"},
		{name: "review get", method: http.MethodGet, path: "/api/reviews/123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := testutil.PerformRequest(r, tt.method, tt.path, nil)
			if resp.Code == http.StatusUnauthorized {
				t.Fatalf("expected route to be public, got %d body=%s", resp.Code, resp.Body.String())
			}
		})
	}
}

func TestProtectedRouteAllowsBearerToken(t *testing.T) {
	r := setupRouterTest(t)

	req := httptest.NewRequest(http.MethodPut, "/api/users/123", nil)
	req.Header.Set("Authorization", "Bearer test-token")

	resp := testutil.PerformRawRequest(r, req)
	if resp.Code == http.StatusUnauthorized {
		t.Fatalf("expected authenticated request to pass middleware, got %d body=%s", resp.Code, resp.Body.String())
	}
}
