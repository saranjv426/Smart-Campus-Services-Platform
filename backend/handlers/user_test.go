package handlers

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func TestGetUserReturnsUser(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)

	handler := NewUserHandler(db)
	router := gin.New()
	router.GET("/users/:id", handler.GetUser)

	rec := performRequest(t, router, http.MethodGet, "/users/"+user.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[models.User](t, rec)
	if resp.ID != user.ID {
		t.Fatalf("expected user ID %s, got %s", user.ID, resp.ID)
	}
	if resp.Email != user.Email {
		t.Fatalf("expected user email %s, got %s", user.Email, resp.Email)
	}
}

func TestGetUserReturnsNotFoundForMissingUser(t *testing.T) {
	db := setupTestDB(t)

	handler := NewUserHandler(db)
	router := gin.New()
	router.GET("/users/:id", handler.GetUser)

	rec := performRequest(t, router, http.MethodGet, "/users/missing", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestUpdateUserPersistsChanges(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)

	handler := NewUserHandler(db)
	router := gin.New()
	router.PUT("/users/:id", handler.UpdateUser)

	rec := performRequest(t, router, http.MethodPut, "/users/"+user.ID, UpdateUserRequest{
		FirstName:  "Updated",
		LastName:   "Name",
		Phone:      "555-9999",
		Department: "Engineering",
		AvatarURL:  "https://example.com/avatar.png",
		Bio:        "Updated bio",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[models.User](t, rec)
	if resp.FirstName != "Updated" || resp.Department != "Engineering" {
		t.Fatalf("expected updated user in response, got %+v", resp)
	}

	var updated models.User
	if err := db.First(&updated, "id = ?", user.ID).Error; err != nil {
		t.Fatalf("failed to reload user: %v", err)
	}
	if updated.Department != "Engineering" {
		t.Fatalf("expected department to be updated, got %s", updated.Department)
	}
}

func TestUpdateUserReturnsBadRequestForInvalidPayload(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)

	handler := NewUserHandler(db)
	router := gin.New()
	router.PUT("/users/:id", handler.UpdateUser)

	req, err := http.NewRequest(http.MethodPut, "/users/"+user.ID, bytes.NewBufferString("{"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := performRawRequest(router, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestGetProfileReturnsBookingsAndReviews(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID)
	createReviewFixture(t, db, user.ID, service.ID, 5)

	handler := NewUserHandler(db)
	router := gin.New()
	router.GET("/users/:id/profile", handler.GetProfile)

	rec := performRequest(t, router, http.MethodGet, "/users/"+user.ID+"/profile", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	profile := decodeJSON[models.User](t, rec)
	if len(profile.Bookings) != 1 || len(profile.Reviews) != 1 {
		t.Fatalf("expected one booking and one review, got %+v", profile)
	}
	if profile.Bookings[0].UserID != user.ID || profile.Reviews[0].UserID != user.ID {
		t.Fatalf("expected preloaded associations for user %s, got %+v", user.ID, profile)
	}
}

func TestGetProfileReturnsNotFoundForMissingUser(t *testing.T) {
	db := setupTestDB(t)

	handler := NewUserHandler(db)
	router := gin.New()
	router.GET("/users/:id/profile", handler.GetProfile)

	rec := performRequest(t, router, http.MethodGet, "/users/missing/profile", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}
}
