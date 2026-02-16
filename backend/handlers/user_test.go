package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"smart-campus-services/models"
	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
)

func setupUserRouter(t *testing.T) (*gin.Engine, string) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	h := NewUserHandler(db)

	user := models.User{
		Email:     "user1@campus.edu",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "+1111111111",
		Role:      "student",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to seed user: %v", err)
	}

	r := gin.New()
	r.GET("/api/users/:id", h.GetUser)
	r.PUT("/api/users/:id", h.UpdateUser)
	r.GET("/api/users/:id/profile", h.GetProfile)

	return r, user.ID
}

func TestGetUserSuccess(t *testing.T) {
	r, userID := setupUserRouter(t)

	resp := testutil.PerformRequest(r, http.MethodGet, "/api/users/"+userID, nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	var data map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &data); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if data["id"] != userID {
		t.Fatalf("expected id=%s, got %v", userID, data["id"])
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	r, userID := setupUserRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"firstName":  "Jane",
		"department": "Computer Science",
		"bio":        "Student profile",
	})

	resp := testutil.PerformRequest(r, http.MethodPut, "/api/users/"+userID, body)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	check := testutil.PerformRequest(r, http.MethodGet, "/api/users/"+userID, nil)
	if check.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, check.Code, check.Body.String())
	}

	var updated map[string]any
	if err := json.Unmarshal(check.Body.Bytes(), &updated); err != nil {
		t.Fatalf("failed to decode updated user: %v", err)
	}
	if updated["firstName"] != "Jane" {
		t.Fatalf("expected firstName=Jane, got %v", updated["firstName"])
	}
}

func TestGetProfileUserNotFound(t *testing.T) {
	r, _ := setupUserRouter(t)

	resp := testutil.PerformRequest(r, http.MethodGet, "/api/users/non-existent-id/profile", nil)
	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected %d, got %d body=%s", http.StatusNotFound, resp.Code, resp.Body.String())
	}
}

func TestGetProfileSuccess(t *testing.T) {
	r, userID := setupUserRouter(t)

	resp := testutil.PerformRequest(r, http.MethodGet, "/api/users/"+userID+"/profile", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	var profile map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &profile); err != nil {
		t.Fatalf("failed to decode profile: %v", err)
	}
	if profile["id"] != userID {
		t.Fatalf("expected id=%s, got %v", userID, profile["id"])
	}
}

func TestUpdateUserInvalidJSON(t *testing.T) {
	r, userID := setupUserRouter(t)

	resp := testutil.PerformRequest(r, http.MethodPut, "/api/users/"+userID, []byte("{invalid json"))
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}

func TestGetUserNotFound(t *testing.T) {
	r, _ := setupUserRouter(t)

	resp := testutil.PerformRequest(r, http.MethodGet, "/api/users/missing", nil)
	if resp.Code != http.StatusNotFound {
		t.Fatalf("expected %d, got %d body=%s", http.StatusNotFound, resp.Code, resp.Body.String())
	}

	var data map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &data); err != nil {
		t.Fatalf("failed to decode error: %v", err)
	}
	if data["error"] != "User not found" {
		t.Fatalf("unexpected error response: %s", fmt.Sprint(data["error"]))
	}
}
