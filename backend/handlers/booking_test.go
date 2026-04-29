package handlers

import (
	"bytes"
	"net/http"
	"strings"
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
	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)

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
	if booking.UserID != user.ID || booking.ServiceID != service.ID || booking.Notes != "Need transportation" {
		t.Fatalf("expected response body to include created booking details, got %+v", booking)
	}
}

func TestCreateBookingReturnsBadRequestForMissingRequiredFields(t *testing.T) {
	db := setupTestDB(t)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)

	rec := performRequest(t, router, http.MethodPost, "/bookings", map[string]any{
		"notes": "Missing user, service, and times",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[map[string]string](t, rec)
	if resp["error"] == "" {
		t.Fatalf("expected error response body, got %+v", resp)
	}
}

func TestCreateBookingRejectsMalformedJSON(t *testing.T) {
	db := setupTestDB(t)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)

	req, err := http.NewRequest(http.MethodPost, "/bookings", bytes.NewBufferString("{"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := performRawRequest(router, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestGetAllBookings(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID)
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "approved"
	})

	handler := NewBookingHandler(db)
	router := gin.New()
	router.GET("/bookings", handler.GetAllBookings)

	rec := performRequest(t, router, http.MethodGet, "/bookings", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 2 {
		t.Fatalf("expected 2 bookings, got %+v", bookings)
	}
	for _, booking := range bookings {
		if booking.User.ID != user.ID || booking.Service.ID != service.ID {
			t.Fatalf("expected user and service to be preloaded, got %+v", booking)
		}
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
	router := gin.New()
	router.GET("/bookings/:id", handler.GetBooking)

	rec := performRequest(t, router, http.MethodGet, "/bookings/missing", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[map[string]string](t, rec)
	if resp["error"] != "Booking not found" {
		t.Fatalf("expected booking not found error, got %+v", resp)
	}
}

func TestGetUserBookingsWithoutStatusFilterReturnsAllUserRecords(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	otherUser := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID)
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "approved"
	})
	createBookingFixture(t, db, otherUser.ID, service.ID)

	handler := NewBookingHandler(db)
	router := setupBookingTestRouter(handler)

	rec := performRequest(t, router, http.MethodGet, "/bookings/user/"+user.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 2 {
		t.Fatalf("expected two bookings for user %s, got %+v", user.ID, bookings)
	}

	for _, booking := range bookings {
		if booking.UserID != user.ID {
			t.Fatalf("expected booking to belong to user %s, got %+v", user.ID, bookings)
		}
		if booking.Service.ID != service.ID {
			t.Fatalf("expected service to be preloaded for booking %+v", booking)
		}
	}
}

func TestGetUserBookingsWithValidStatusFilterReturnsMatchingRecords(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "pending"
	})
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "approved"
	})
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "rejected"
	})

	handler := NewBookingHandler(db)
	router := gin.New()
	router.GET("/bookings/user/:userId", handler.GetUserBookings)

	rec := performRequest(t, router, http.MethodGet, "/bookings/user/"+user.ID+"?status=approved", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 1 {
		t.Fatalf("expected one approved booking, got %+v", bookings)
	}
	if bookings[0].Status != "approved" {
		t.Fatalf("expected approved booking, got %+v", bookings[0])
	}
	if bookings[0].Service.ID != service.ID {
		t.Fatalf("expected service to be preloaded for booking %+v", bookings[0])
	}
}

func TestGetUserBookingsWithInvalidStatusFilterReturnsBadRequest(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.GET("/bookings/user/:userId", handler.GetUserBookings)

	rec := performRequest(t, router, http.MethodGet, "/bookings/user/"+user.ID+"?status=unknown", nil)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[map[string]string](t, rec)
	if !strings.Contains(strings.ToLower(resp["error"]), "invalid status") {
		t.Fatalf("expected clear invalid status error, got %+v", resp)
	}
}

func TestGetUserBookingsWithNoMatchingStatusReturnsEmptyArray(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "pending"
	})

	handler := NewBookingHandler(db)
	router := gin.New()
	router.GET("/bookings/user/:userId", handler.GetUserBookings)

	rec := performRequest(t, router, http.MethodGet, "/bookings/user/"+user.ID+"?status=cancelled", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 0 {
		t.Fatalf("expected no matching bookings, got %+v", bookings)
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

	resp := decodeJSON[models.Booking](t, rec)
	if resp.ID != booking.ID || resp.Notes != "Updated note" || resp.Status != "pending" {
		t.Fatalf("expected response body to include updated booking, got %+v", resp)
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "cancelled" {
		t.Fatalf("expected status cancelled, got %s", updated.Status)
	}
	if updated.Status != "pending" {
		t.Fatalf("expected update to preserve booking status, got %s", updated.Status)
	}
}

func TestUpdateBookingReturnsNotFoundForMissingBooking(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.PUT("/bookings/:id", handler.UpdateBooking)

	now := time.Now().UTC()
	rec := performRequest(t, router, http.MethodPut, "/bookings/missing", CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: now.Add(4 * time.Hour),
		EndTime:   now.Add(5 * time.Hour),
		Notes:     "Updated note",
	})

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[map[string]string](t, rec)
	if resp["error"] != "Booking not found" {
		t.Fatalf("expected booking not found error, got %+v", resp)
	}
}

func TestUpdateBookingReturnsBadRequestForInvalidPayload(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.PUT("/bookings/:id", handler.UpdateBooking)

	rec := performRequest(t, router, http.MethodPut, "/bookings/"+booking.ID, map[string]any{
		"notes": "Missing required fields",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestCancelBooking(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.DELETE("/bookings/:id", handler.CancelBooking)

	rec := performRequest(t, router, http.MethodDelete, "/bookings/"+booking.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "cancelled" {
		t.Fatalf("expected booking status cancelled, got %s", updated.Status)
	}
}

func TestCancelBookingReturnsNotFoundForMissingBooking(t *testing.T) {
	db := setupTestDB(t)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.DELETE("/bookings/:id", handler.CancelBooking)

	rec := performRequest(t, router, http.MethodDelete, "/bookings/missing", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[map[string]string](t, rec)
	if resp["error"] != "Booking not found" {
		t.Fatalf("expected booking not found error, got %+v", resp)
	}
}
