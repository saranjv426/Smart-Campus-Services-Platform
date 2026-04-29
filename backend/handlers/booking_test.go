package handlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func setupBookingTestRouter(handler *BookingHandler) *gin.Engine {
	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)
	router.GET("/bookings/:id", handler.GetBooking)
	router.GET("/bookings/user/:userId", handler.GetUserBookings)
	router.PUT("/bookings/:id", handler.UpdateBooking)
	router.DELETE("/bookings/:id", handler.CancelBooking)
	router.PATCH("/bookings/:id/status", handler.CancelBooking)
	return router
}

func TestCreateBookingSetsPendingStatus(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	now := time.Now().UTC()
	rec := performRequest(t, router, http.MethodPost, "/bookings", CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: now.Add(time.Hour),
		EndTime:   now.Add(2 * time.Hour),
		Notes:     "Need transportation",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	booking := decodeJSON[models.Booking](t, rec)
	if booking.Status != "pending" {
		t.Fatalf("expected booking status pending, got %s", booking.Status)
	}
}

func TestCreateBookingRequiresEndTime(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/bookings", map[string]any{
		"userId":    user.ID,
		"serviceId": service.ID,
		"startTime": time.Now().UTC().Add(time.Hour),
		"notes":     "Missing end time",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestCreateBookingRejectsEndTimeBeforeStartTime(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	start := time.Now().UTC().Add(2 * time.Hour)
	rec := performRequest(t, router, http.MethodPost, "/bookings", CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: start,
		EndTime:   start.Add(-time.Hour),
		Notes:     "Invalid time range",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestCreateBookingCreatesNotification(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	start := time.Now().UTC().Add(2 * time.Hour)
	rec := performRequest(t, router, http.MethodPost, "/bookings", CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: start,
		EndTime:   start.Add(time.Hour),
		Notes:     "Need slot",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
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

func TestGetBookingReturnsBookingWithRelations(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodGet, "/bookings/"+booking.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[models.Booking](t, rec)
	if resp.User.ID != user.ID || resp.Service.ID != service.ID {
		t.Fatalf("expected preloaded user and service, got %+v", resp)
	}
}

func TestGetBookingNotFound(t *testing.T) {
	db := setupTestDB(t)
	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodGet, "/bookings/missing", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestGetUserBookingsReturnsUserRecords(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	otherUser := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID)
	createBookingFixture(t, db, otherUser.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodGet, "/bookings/user/"+user.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 1 || bookings[0].UserID != user.ID {
		t.Fatalf("expected one booking for user %s, got %+v", user.ID, bookings)
	}
}

func TestCancelBookingDeleteMarksStatusCancelled(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodDelete, "/bookings/"+booking.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "cancelled" {
		t.Fatalf("expected status cancelled, got %s", updated.Status)
	}
}

func TestUpdateBookingPersistsChanges(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	now := time.Now().UTC()
	rec := performRequest(t, router, http.MethodPut, "/bookings/"+booking.ID, CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: now.Add(4 * time.Hour),
		EndTime:   now.Add(5 * time.Hour),
		Notes:     "Updated note",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Notes != "Updated note" {
		t.Fatalf("expected booking notes to be updated, got %s", updated.Notes)
	}
}

func TestUpdateBookingRejectsEndTimeBeforeStartTime(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	start := time.Now().UTC().Add(4 * time.Hour)
	rec := performRequest(t, router, http.MethodPut, "/bookings/"+booking.ID, CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: start,
		EndTime:   start.Add(-time.Hour),
		Notes:     "Invalid update",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestCancelBookingStatusSuccess(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodPatch, "/bookings/"+booking.ID+"/status", UpdateBookingStatusRequest{
		Status: "cancelled",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
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
	db := setupTestDB(t)
	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodPatch, "/bookings/booking-1/status", UpdateBookingStatusRequest{
		Status: "completed",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestCancelBookingReturnsNotFoundForMissingBooking(t *testing.T) {
	db := setupTestDB(t)
	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodPatch, "/bookings/missing/status", UpdateBookingStatusRequest{
		Status: "cancelled",
	})

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}
}
