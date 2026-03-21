package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"smart-campus-services/testutil"
	"smart-campus-services/validation"

	"github.com/gin-gonic/gin"
)

func setupAuthRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := testutil.NewTestDB(t)
	h := NewAuthHandler(db)

	r := gin.New()
	r.POST("/api/auth/register", h.Register)
	r.POST("/api/auth/login", h.Login)
	r.POST("/api/auth/logout", h.Logout)
	r.POST("/api/auth/refresh", h.RefreshToken)
	return r
}

func TestRegisterSuccess(t *testing.T) {
	r := setupAuthRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"email":     "student1@uf.edu",
		"password":  "password123",
		"firstName": "John",
		"lastName":  "Doe",
		"phone":     "+1234567890",
		"role":      "student",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/auth/register", body)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusCreated, resp.Code, resp.Body.String())
	}

	var data map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &data); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if data["role"] != "student" {
		t.Fatalf("expected role=student, got %v", data["role"])
	}
	if data["serviceId"] == nil {
		t.Fatalf("expected serviceId key in response")
	}
}

func TestRegisterRejectsInvalidRole(t *testing.T) {
	r := setupAuthRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"email":     "student2@uf.edu",
		"password":  "password123",
		"firstName": "Jane",
		"lastName":  "Doe",
		"phone":     "+1234567890",
		"role":      "manager",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/auth/register", body)
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}

func TestLoginSuccess(t *testing.T) {
	r := setupAuthRouter(t)

	registerBody := testutil.MustMarshal(map[string]any{
		"email":     "student3@uf.edu",
		"password":  "password123",
		"firstName": "A",
		"lastName":  "B",
		"phone":     "+1234567890",
		"role":      "student",
	})
	registerResp := testutil.PerformRequest(r, http.MethodPost, "/api/auth/register", registerBody)
	if registerResp.Code != http.StatusCreated {
		t.Fatalf("register failed: status=%d body=%s", registerResp.Code, registerResp.Body.String())
	}

	loginBody := testutil.MustMarshal(map[string]any{
		"email":    "student3@uf.edu",
		"password": "password123",
	})
	loginResp := testutil.PerformRequest(r, http.MethodPost, "/api/auth/login", loginBody)
	if loginResp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusOK, loginResp.Code, loginResp.Body.String())
	}
}
