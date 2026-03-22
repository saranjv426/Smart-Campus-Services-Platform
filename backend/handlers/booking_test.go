package handlers

import (
	"net/http"
	"testing"
	"time"

<<<<<<< HEAD
	"smart-campus-services/models"
	"smart-campus-services/testutil"

=======
>>>>>>> f85f4d5 (Add Sprint 2 backend tests and documentation)
	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func TestCreateBookingSetsPendingStatus(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)

<<<<<<< HEAD
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
=======
	now := time.Now().UTC()
	rec := performRequest(t, router, http.MethodPost, "/bookings", CreateBookingRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		StartTime: now.Add(time.Hour),
		EndTime:   now.Add(2 * time.Hour),
		Notes:     "Need transportation",
>>>>>>> f85f4d5 (Add Sprint 2 backend tests and documentation)
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	booking := decodeJSON[models.Booking](t, rec)
	if booking.Status != "pending" {
		t.Fatalf("expected booking status pending, got %s", booking.Status)
	}
}

func TestGetBookingReturnsBookingWithRelations(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.GET("/bookings/:id", handler.GetBooking)

	rec := performRequest(t, router, http.MethodGet, "/bookings/"+booking.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[models.Booking](t, rec)
	if resp.User.ID != user.ID || resp.Service.ID != service.ID {
		t.Fatalf("expected preloaded user and service, got %+v", resp)
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
	router := gin.New()
	router.GET("/bookings/user/:userId", handler.GetUserBookings)

	rec := performRequest(t, router, http.MethodGet, "/bookings/user/"+user.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 1 || bookings[0].UserID != user.ID {
		t.Fatalf("expected one booking for user %s, got %+v", user.ID, bookings)
	}
}

func TestUpdateBookingPersistsChanges(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewBookingHandler(db)
	router := gin.New()
	router.PUT("/bookings/:id", handler.UpdateBooking)

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

func TestCancelBookingMarksStatusCancelled(t *testing.T) {
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
