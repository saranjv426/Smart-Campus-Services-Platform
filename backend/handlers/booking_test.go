package handlers

import (
	"net/http"
	"testing"
	"time"

	"smart-campus-services/models"
	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
)

func setupBookingRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	h := NewBookingHandler(db)

	r := gin.New()
	r.POST("/api/bookings", h.CreateBooking)
	r.PATCH("/api/bookings/:id/status", h.CancelBooking)
	return r
}

func TestCreateBookingValidation(t *testing.T) {
	r := setupBookingRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"userId":    "user-1",
		"serviceId": "service-1",
		"startTime": time.Now().UTC(),
		"notes":     "Need slot",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/bookings", body)
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}

func TestCreateBookingSuccess(t *testing.T) {
	r := setupBookingRouter(t)
	start := time.Now().UTC().Add(2 * time.Hour)
	end := start.Add(time.Hour)

	body := testutil.MustMarshal(map[string]any{
		"userId":    "user-1",
		"serviceId": "service-1",
		"startTime": start,
		"endTime":   end,
		"notes":     "Need slot",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/bookings", body)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusCreated, resp.Code, resp.Body.String())
	}
}

func TestCreateBookingCreatesNotification(t *testing.T) {
	db := testutil.NewTestDB(t)
	h := NewBookingHandler(db)

	user := models.User{
		Email:     "booking-notif@campus.edu",
		Password:  "password123",
		FirstName: "Booking",
		LastName:  "User",
		Phone:     "+3333333333",
		Role:      "student",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to seed user: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/api/bookings", h.CreateBooking)

	start := time.Now().UTC().Add(2 * time.Hour)
	end := start.Add(time.Hour)
	body := testutil.MustMarshal(map[string]any{
		"userId":    user.ID,
		"serviceId": "service-1",
		"startTime": start,
		"endTime":   end,
		"notes":     "Need slot",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/bookings", body)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusCreated, resp.Code, resp.Body.String())
	}

	var notifications []models.Notification
	if err := db.Where("user_id = ?", user.ID).Find(&notifications).Error; err != nil {
		t.Fatalf("failed to fetch notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
	if notifications[0].Title != "Booking Request Submitted" {
		t.Fatalf("expected booking notification title, got %s", notifications[0].Title)
	}
	if notifications[0].IsRead {
		t.Fatalf("expected notification to be unread")
	}
}

func TestCancelBookingStatusSuccess(t *testing.T) {
	db := testutil.NewTestDB(t)
	h := NewBookingHandler(db)

	booking := models.Booking{
		UserID:    "user-1",
		ServiceID: "service-1",
		Status:    "pending",
		StartTime: time.Now().UTC().Add(2 * time.Hour),
		EndTime:   time.Now().UTC().Add(3 * time.Hour),
	}
	if err := db.Create(&booking).Error; err != nil {
		t.Fatalf("failed to create booking: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.PATCH("/api/bookings/:id/status", h.CancelBooking)

	req := testutil.NewJSONRequest(t, http.MethodPatch, "/api/bookings/"+booking.ID+"/status", testutil.MustMarshal(map[string]any{
		"status": "cancelled",
	}))
	resp := testutil.PerformRawRequest(r, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "cancelled" {
		t.Fatalf("expected status cancelled, got %s", updated.Status)
	}
}

func TestCancelBookingStatusRejectsInvalidStatus(t *testing.T) {
	r := setupBookingRouter(t)

	req := testutil.NewJSONRequest(t, http.MethodPatch, "/api/bookings/booking-1/status", testutil.MustMarshal(map[string]any{
		"status": "completed",
	}))
	resp := testutil.PerformRawRequest(r, req)
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}
