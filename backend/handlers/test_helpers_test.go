package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"smart-campus-services/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	gin.SetMode(gin.TestMode)

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", uuid.NewString())
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Booking{},
		&models.Notification{},
		&models.Review{},
	); err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}

func performRequest(t *testing.T, router *gin.Engine, method, path string, body any) *httptest.ResponseRecorder {
	t.Helper()

	var reader *bytes.Reader
	if body == nil {
		reader = bytes.NewReader(nil)
	} else {
		payload, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("failed to marshal request body: %v", err)
		}
		reader = bytes.NewReader(payload)
	}

	req, err := http.NewRequest(method, path, reader)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func performRawRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func decodeJSON[T any](t *testing.T, rec *httptest.ResponseRecorder) T {
	t.Helper()

	var payload T
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	return payload
}

func createServiceFixture(t *testing.T, db *gorm.DB, overrides ...func(*models.Service)) models.Service {
	t.Helper()

	service := models.Service{
		Name:        "Campus Shuttle",
		Description: "Transportation support",
		Category:    "transportation",
		Location:    "Transit Hub",
		Email:       "transport@campus.edu",
		Hours:       "Daily",
		IsActive:    true,
	}

	for _, override := range overrides {
		override(&service)
	}

	if err := db.Create(&service).Error; err != nil {
		t.Fatalf("failed to create service fixture: %v", err)
	}

	return service
}

func createUserFixture(t *testing.T, db *gorm.DB, overrides ...func(*models.User)) models.User {
	t.Helper()

	user := models.User{
		Email:     fmt.Sprintf("%s@campus.edu", uuid.NewString()),
		Password:  "password123",
		FirstName: "Test",
		LastName:  "User",
		Phone:     "555-0100",
		Role:      "student",
	}

	for _, override := range overrides {
		override(&user)
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to create user fixture: %v", err)
	}

	return user
}

func createBookingFixture(t *testing.T, db *gorm.DB, userID, serviceID string, overrides ...func(*models.Booking)) models.Booking {
	t.Helper()

	now := time.Now().UTC()
	booking := models.Booking{
		UserID:    userID,
		ServiceID: serviceID,
		Status:    "pending",
		StartTime: now.Add(2 * time.Hour),
		EndTime:   now.Add(3 * time.Hour),
		Notes:     "Initial booking",
	}

	for _, override := range overrides {
		override(&booking)
	}

	if err := db.Create(&booking).Error; err != nil {
		t.Fatalf("failed to create booking fixture: %v", err)
	}

	return booking
}

func createNotificationFixture(t *testing.T, db *gorm.DB, userID string, overrides ...func(*models.Notification)) models.Notification {
	t.Helper()

	notification := models.Notification{
		UserID:  userID,
		Title:   "Reminder",
		Message: "Your booking is tomorrow",
		Type:    "reminder",
		IsRead:  false,
	}

	for _, override := range overrides {
		override(&notification)
	}

	if err := db.Create(&notification).Error; err != nil {
		t.Fatalf("failed to create notification fixture: %v", err)
	}

	return notification
}

func createReviewFixture(t *testing.T, db *gorm.DB, userID, serviceID string, rating int, overrides ...func(*models.Review)) models.Review {
	t.Helper()

	review := models.Review{
		UserID:    userID,
		ServiceID: serviceID,
		Rating:    rating,
		Comment:   "Helpful service",
	}

	for _, override := range overrides {
		override(&review)
	}

	if err := db.Create(&review).Error; err != nil {
		t.Fatalf("failed to create review fixture: %v", err)
	}

	return review
}
