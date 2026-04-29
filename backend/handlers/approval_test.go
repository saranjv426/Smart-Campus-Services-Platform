package handlers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
	"smart-campus-services/testutil"
)

func setupApprovalRouter(handler *ApprovalHandler) *gin.Engine {
	router := gin.New()
	router.GET("/approval/staff/:staffId/pending", handler.GetPendingBookings)
	router.PUT("/approval/bookings/:id/approve", handler.ApproveBooking)
	return router
}

func TestGetPendingBookingsReturnsStaffServiceBookings(t *testing.T) {
	db := testutil.NewTestDB(t)
	service := createServiceFixture(t, db)
	otherService := createServiceFixture(t, db, func(s *models.Service) {
		s.Name = "Health Center"
		s.Category = "health"
	})
	staff := createUserFixture(t, db, func(u *models.User) {
		u.Role = "staff"
		u.ServiceID = service.ID
	})
	user := createUserFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "pending"
	})
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "approved"
	})
	createBookingFixture(t, db, user.ID, otherService.ID, func(b *models.Booking) {
		b.Status = "pending"
	})

	handler := NewApprovalHandler(db)
	router := setupApprovalRouter(handler)

	rec := testutil.PerformRequest(router, http.MethodGet, "/approval/staff/"+staff.ID+"/pending", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 1 || bookings[0].ServiceID != service.ID {
		t.Fatalf("expected one pending booking for service %s, got %+v", service.ID, bookings)
	}
}

func TestGetAllBookingsForServiceReturnsAllStatuses(t *testing.T) {
	db := setupTestDB(t)
	service := createServiceFixture(t, db)
	staff := createUserFixture(t, db, func(u *models.User) {
		u.Role = "staff"
		u.ServiceID = service.ID
	})
	user := createUserFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "pending"
	})
	createBookingFixture(t, db, user.ID, service.ID, func(b *models.Booking) {
		b.Status = "approved"
	})

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.GET("/approval/staff/:staffId/all", handler.GetAllBookingsForService)

	rec := performRequest(t, router, http.MethodGet, "/approval/staff/"+staff.ID+"/all", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 2 {
		t.Fatalf("expected 2 bookings, got %d", len(bookings))
	}
}

func TestApproveBookingUpdatesStatusAndCreatesNotification(t *testing.T) {
	db := testutil.NewTestDB(t)
	service := createServiceFixture(t, db)
	staff := createUserFixture(t, db, func(u *models.User) {
		u.Role = "staff"
		u.ServiceID = service.ID
	})
	user := createUserFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewApprovalHandler(db)
	router := setupApprovalRouter(handler)

	rec := testutil.PerformRequest(router, http.MethodPut, "/approval/bookings/"+booking.ID+"/approve", testutil.MustMarshal(ApprovalRequest{
		Status:        "approved",
		ApprovalNotes: "See you soon",
		StaffID:       staff.ID,
	}))

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "approved" || updated.ApprovedBy != staff.ID {
		t.Fatalf("expected approved booking with approver %s, got %+v", staff.ID, updated)
	}

	var notifications []models.Notification
	if err := db.Where("user_id = ?", user.ID).Find(&notifications).Error; err != nil {
		t.Fatalf("failed to load notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
	if notifications[0].Type != "booking_approval" || notifications[0].Title != "Booking Approved" {
		t.Fatalf("unexpected notification stored: %+v", notifications[0])
	}
	if !strings.Contains(notifications[0].Message, "See you soon") {
		t.Fatalf("expected notification message to include approval notes, got %q", notifications[0].Message)
	}
}

func TestRejectBookingCreatesNotification(t *testing.T) {
	db := setupTestDB(t)
	service := createServiceFixture(t, db)
	staff := createUserFixture(t, db, func(u *models.User) {
		u.Role = "staff"
		u.ServiceID = service.ID
	})
	user := createUserFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.PUT("/approval/bookings/:id/reject", handler.RejectBooking)

	rec := performRequest(t, router, http.MethodPut, "/approval/bookings/"+booking.ID+"/reject", ApprovalRequest{
		Status:        "rejected",
		ApprovalNotes: "Missing details",
		StaffID:       staff.ID,
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var notifications []models.Notification
	if err := db.Where("user_id = ?", user.ID).Find(&notifications).Error; err != nil {
		t.Fatalf("failed to load notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
	if notifications[0].Type != "booking_rejection" || notifications[0].Title != "Booking Rejected" {
		t.Fatalf("unexpected notification stored: %+v", notifications[0])
	}
	if !strings.Contains(notifications[0].Message, "Missing details") {
		t.Fatalf("expected notification message to include rejection reason, got %q", notifications[0].Message)
	}
}

func TestRejectBookingRequiresMatchingStaffService(t *testing.T) {
	db := setupTestDB(t)
	service := createServiceFixture(t, db)
	otherService := createServiceFixture(t, db, func(s *models.Service) {
		s.Name = "Dining Hall"
		s.Category = "dining"
	})
	staff := createUserFixture(t, db, func(u *models.User) {
		u.Role = "staff"
		u.ServiceID = otherService.ID
	})
	user := createUserFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.PUT("/approval/bookings/:id/reject", handler.RejectBooking)

	rec := performRequest(t, router, http.MethodPut, "/approval/bookings/"+booking.ID+"/reject", ApprovalRequest{
		Status:        "rejected",
		ApprovalNotes: "Wrong service owner",
		StaffID:       staff.ID,
	})

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestGetAllPendingBookingsRequiresAdmin(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.GET("/approval/admin/:userId/pending", handler.GetAllPendingBookings)

	rec := performRequest(t, router, http.MethodGet, "/approval/admin/"+user.ID+"/pending", nil)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestGetAllBookingsReturnsAdminView(t *testing.T) {
	db := setupTestDB(t)
	admin := createUserFixture(t, db, func(u *models.User) {
		u.Role = "admin"
	})
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createBookingFixture(t, db, user.ID, service.ID)

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.GET("/approval/admin/:userId/all", handler.GetAllBookings)

	rec := performRequest(t, router, http.MethodGet, "/approval/admin/"+admin.ID+"/all", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	bookings := decodeJSON[[]models.Booking](t, rec)
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking, got %d", len(bookings))
	}
}

func TestAdminApproveBookingUpdatesStatus(t *testing.T) {
	db := setupTestDB(t)
	admin := createUserFixture(t, db, func(u *models.User) {
		u.Role = "admin"
	})
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.PUT("/approval/admin/:userId/bookings/:id/approve", handler.AdminApproveBooking)

	rec := performRequest(t, router, http.MethodPut, "/approval/admin/"+admin.ID+"/bookings/"+booking.ID+"/approve", ApprovalRequest{
		Status:        "approved",
		ApprovalNotes: "Admin approved",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "approved" || updated.ApprovedBy != admin.ID {
		t.Fatalf("expected admin-approved booking, got %+v", updated)
	}

	var notifications []models.Notification
	if err := db.Where("user_id = ?", user.ID).Find(&notifications).Error; err != nil {
		t.Fatalf("failed to load notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
	if notifications[0].Type != "booking_approval" || notifications[0].Title != "Booking Approved" {
		t.Fatalf("unexpected notification stored: %+v", notifications[0])
	}
	if !strings.Contains(notifications[0].Message, "approved by admin") {
		t.Fatalf("expected admin approval wording in notification, got %q", notifications[0].Message)
	}
}

func TestAdminRejectBookingUpdatesStatus(t *testing.T) {
	db := setupTestDB(t)
	admin := createUserFixture(t, db, func(u *models.User) {
		u.Role = "admin"
	})
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	booking := createBookingFixture(t, db, user.ID, service.ID)

	handler := NewApprovalHandler(db)
	router := gin.New()
	router.PUT("/approval/admin/:userId/bookings/:id/reject", handler.AdminRejectBooking)

	rec := performRequest(t, router, http.MethodPut, "/approval/admin/"+admin.ID+"/bookings/"+booking.ID+"/reject", ApprovalRequest{
		Status:        "rejected",
		ApprovalNotes: "Admin rejected",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Booking
	if err := db.First(&updated, "id = ?", booking.ID).Error; err != nil {
		t.Fatalf("failed to reload booking: %v", err)
	}
	if updated.Status != "rejected" || updated.ApprovedBy != admin.ID {
		t.Fatalf("expected admin-rejected booking, got %+v", updated)
	}

	var notifications []models.Notification
	if err := db.Where("user_id = ?", user.ID).Find(&notifications).Error; err != nil {
		t.Fatalf("failed to load notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
	if notifications[0].Type != "booking_rejection" || notifications[0].Title != "Booking Rejected" {
		t.Fatalf("unexpected notification stored: %+v", notifications[0])
	}
	if !strings.Contains(notifications[0].Message, "Admin rejected") {
		t.Fatalf("expected rejection reason in notification, got %q", notifications[0].Message)
	}
}
