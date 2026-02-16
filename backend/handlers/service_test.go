package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
)

func setupServiceRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	h := NewServiceHandler(db)

	r := gin.New()
	r.POST("/api/services", h.CreateService)
	r.GET("/api/services", h.ListServices)
	return r
}

func TestCreateServiceValidation(t *testing.T) {
	r := setupServiceRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"description": "Service description",
		"category":    "dining",
		"location":    "Building A",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/services", body)
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}

func TestCreateAndListService(t *testing.T) {
	r := setupServiceRouter(t)

	createBody := testutil.MustMarshal(map[string]any{
		"name":        "Campus Clinic",
		"description": "Health support for students",
		"category":    "health",
		"location":    "Center Building",
		"phone":       "+1111111111",
		"email":       "clinic@campus.edu",
		"hours":       "9-5",
	})

	createResp := testutil.PerformRequest(r, http.MethodPost, "/api/services", createBody)
	if createResp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusCreated, createResp.Code, createResp.Body.String())
	}

	listResp := testutil.PerformRequest(r, http.MethodGet, "/api/services", nil)
	if listResp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusOK, listResp.Code, listResp.Body.String())
	}

	var services []map[string]any
	if err := json.Unmarshal(listResp.Body.Bytes(), &services); err != nil {
		t.Fatalf("failed to decode services: %v", err)
	}
	if len(services) != 1 {
		t.Fatalf("expected 1 service, got %d", len(services))
	}
}
