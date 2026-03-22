package handlers

import (
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

	var updated models.User
	if err := db.First(&updated, "id = ?", user.ID).Error; err != nil {
		t.Fatalf("failed to reload user: %v", err)
	}
	if updated.Department != "Engineering" {
		t.Fatalf("expected department to be updated, got %s", updated.Department)
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
}
