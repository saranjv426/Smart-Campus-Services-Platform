package middleware

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthRequiredRejectsMissingAuthorizationHeader(t *testing.T) {
	router := gin.New()
	router.Use(AuthRequired())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/protected", "", "")

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestAuthRequiredAllowsBearerTokenAndStoresIt(t *testing.T) {
	router := gin.New()
	router.Use(AuthRequired())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"token": c.GetString("token")})
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/protected", "Authorization", "Bearer demo-token")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d body=%s", rec.Code, rec.Body.String())
	}
	if body := rec.Body.String(); body != "{\"token\":\"demo-token\"}" {
		t.Fatalf("expected token to be stored in context, got %s", body)
	}
}

func TestRequireRolesRejectsForbiddenRole(t *testing.T) {
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("role", "student")
		c.Next()
	})
	router.Use(RequireRoles("admin"))
	router.GET("/admin-only", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/admin-only", "", "")

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestRequireRolesAllowsMatchingRole(t *testing.T) {
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("role", "admin")
		c.Next()
	})
	router.Use(RequireRoles("admin"))
	router.GET("/admin-only", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/admin-only", "", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestCORSSetsHeadersAndShortCircuitsOptions(t *testing.T) {
	router := gin.New()
	router.Use(CORS())
	router.OPTIONS("/resource", func(c *gin.Context) {
		t.Fatal("expected OPTIONS request to be aborted before handler execution")
	})

	rec := performMiddlewareRequest(t, router, http.MethodOptions, "/resource", "", "")

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rec.Code)
	}
	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatalf("expected CORS allow origin header, got %q", rec.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestErrorHandlerWritesInternalServerErrorForUnhandledErrors(t *testing.T) {
	router := gin.New()
	router.Use(ErrorHandler())
	router.GET("/boom", func(c *gin.Context) {
		_ = c.Error(errors.New("boom"))
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/boom", "", "")

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestRequestLoggerAddsRequestIDHeader(t *testing.T) {
	router := gin.New()
	router.Use(RequestLogger())
	router.GET("/logged", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"requestID": c.GetString("requestID")})
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/logged", "", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d body=%s", rec.Code, rec.Body.String())
	}
	if rec.Header().Get(requestIDHeader) == "" {
		t.Fatal("expected request ID header to be present")
	}
}

func TestRequestLoggerPreservesIncomingRequestID(t *testing.T) {
	router := gin.New()
	router.Use(RequestLogger())
	router.GET("/logged", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	rec := performMiddlewareRequest(t, router, http.MethodGet, "/logged", requestIDHeader, "incoming-id")

	if rec.Header().Get(requestIDHeader) != "incoming-id" {
		t.Fatalf("expected existing request ID to be preserved, got %q", rec.Header().Get(requestIDHeader))
	}
}
