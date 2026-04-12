package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"smart-campus-services/middleware"
	"smart-campus-services/models"
	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouterTest(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	r := gin.New()
	RegisterAPIRoutes(r, db)
	return r
}

func setupRouterTestWithDB(t *testing.T) (*gin.Engine, *gorm.DB) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	r := gin.New()
	RegisterAPIRoutes(r, db)
	return r, db
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
		{name: "admin user list", method: http.MethodGet, path: "/api/users"},
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

func TestAdminUsersRouteRequiresAdminRole(t *testing.T) {
	r := setupRouterTest(t)

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	req.Header.Set("Authorization", "Bearer user-2|staff|service-1")

	resp := testutil.PerformRawRequest(r, req)
	if resp.Code != http.StatusForbidden {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusForbidden, resp.Code, resp.Body.String())
	}
}

func TestAdminUsersRouteAllowsAdminRole(t *testing.T) {
	r, db := setupRouterTestWithDB(t)

	admin := models.User{
		ID:        "admin-1",
		Email:     "admin@example.com",
		Password:  "secret123",
		FirstName: "Admin",
		LastName:  "User",
		Phone:     "555-1000",
		Role:      "admin",
	}
	if err := db.Create(&admin).Error; err != nil {
		t.Fatalf("failed to create admin fixture: %v", err)
	}

	user := models.User{
		ID:        "student-1",
		Email:     "student@example.com",
		Password:  "secret123",
		FirstName: "Student",
		LastName:  "User",
		Phone:     "555-2000",
		Role:      "student",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user fixture: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	req.Header.Set("Authorization", "Bearer admin-1|admin|")

	resp := testutil.PerformRawRequest(r, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}
}

func TestAuthRequiredParsesStructuredTokenClaims(t *testing.T) {
	claims, ok := middleware.ParseTokenClaims("user-1|admin|")
	if !ok {
		t.Fatal("expected structured token to parse successfully")
	}
	if claims.UserID != "user-1" || claims.Role != "admin" || claims.ServiceID != "" {
		t.Fatalf("unexpected claims parsed: %+v", claims)
	}
}
