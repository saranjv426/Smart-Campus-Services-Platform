package handlers

import (
	"net/http"
	"time"

	"smart-campus-services/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApprovalHandler struct {
	db *gorm.DB
}

func NewApprovalHandler(db *gorm.DB) *ApprovalHandler {
	return &ApprovalHandler{db: db}
}

type ApprovalRequest struct {
	Status        string `json:"status" binding:"required"` // approved or rejected
	ApprovalNotes string `json:"approvalNotes"`
	StaffID       string `json:"staffId"`
}

// GetPendingBookings returns all pending bookings for a staff member's service
func (h *ApprovalHandler) GetPendingBookings(c *gin.Context) {
	staffID := c.Param("staffId")

	// Get staff member and their service
	var staff models.User
	if err := h.db.First(&staff, "id = ? AND role = ?", staffID, "staff").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff member not found"})
		return
	}

	if staff.ServiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Staff member not assigned to a service"})
		return
	}

	// Get all pending bookings for this service
	var bookings []models.Booking
	if err := h.db.
		Preload("User").
		Preload("Service").
		Where("service_id = ? AND status = ?", staff.ServiceID, "pending").
		Order("created_at DESC").
		Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// GetAllBookingsForService returns all bookings (pending, approved, rejected) for a staff's service
func (h *ApprovalHandler) GetAllBookingsForService(c *gin.Context) {
	staffID := c.Param("staffId")

	// Get staff member and their service
	var staff models.User
	if err := h.db.First(&staff, "id = ? AND role = ?", staffID, "staff").Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff member not found"})
		return
	}

	if staff.ServiceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Staff member not assigned to a service"})
		return
	}

	// Get all bookings for this service
	var bookings []models.Booking
	if err := h.db.
		Preload("User").
		Preload("Service").
		Where("service_id = ?", staff.ServiceID).
		Order("created_at DESC").
		Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// ApproveBooking approves a booking
func (h *ApprovalHandler) ApproveBooking(c *gin.Context) {
	bookingID := c.Param("id")
	var req ApprovalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staffID := c.GetString("staffID") // May be set by middleware
	if staffID == "" {
		staffID = req.StaffID
	}
	if staffID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Staff ID is required"})
		return
	}

	// Get the booking
	var booking models.Booking
	if err := h.db.Preload("Service").First(&booking, "id = ?", bookingID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	// Verify staff manages this service
	var staff models.User
	if err := h.db.First(&staff, "id = ? AND service_id = ?", staffID, booking.ServiceID).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only approve bookings for your service"})
		return
	}

	// Update booking
	updates := map[string]interface{}{
		"status":         "approved",
		"approved_by":    staffID,
		"approval_notes": req.ApprovalNotes,
		"updated_at":     time.Now(),
	}

	if err := h.db.Model(&booking).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve booking"})
		return
	}

	// Create notification for student
	notification := models.Notification{
		UserID:  booking.UserID,
		Title:   "Booking Approved",
		Message: "Your booking for " + booking.Service.Name + " has been approved. Approval notes: " + req.ApprovalNotes,
		Type:    "booking_approval",
		IsRead:  false,
	}
	h.db.Create(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Booking approved successfully", "booking": booking})
}

// RejectBooking rejects a booking
func (h *ApprovalHandler) RejectBooking(c *gin.Context) {
	bookingID := c.Param("id")
	var req ApprovalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	staffID := c.GetString("staffID") // May be set by middleware
	if staffID == "" {
		staffID = req.StaffID
	}
	if staffID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Staff ID is required"})
		return
	}

	// Get the booking
	var booking models.Booking
	if err := h.db.Preload("Service").First(&booking, "id = ?", bookingID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	// Verify staff manages this service
	var staff models.User
	if err := h.db.First(&staff, "id = ? AND service_id = ?", staffID, booking.ServiceID).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only reject bookings for your service"})
		return
	}

	// Update booking
	updates := map[string]interface{}{
		"status":         "rejected",
		"approved_by":    staffID,
		"approval_notes": req.ApprovalNotes,
		"updated_at":     time.Now(),
	}

	if err := h.db.Model(&booking).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject booking"})
		return
	}

	// Create notification for student
	notification := models.Notification{
		UserID:  booking.UserID,
		Title:   "Booking Rejected",
		Message: "Your booking for " + booking.Service.Name + " has been rejected. Reason: " + req.ApprovalNotes,
		Type:    "booking_rejection",
		IsRead:  false,
	}
	h.db.Create(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Booking rejected successfully", "booking": booking})
}

// GetAllPendingBookings returns all pending bookings across all services (for admin)
func (h *ApprovalHandler) GetAllPendingBookings(c *gin.Context) {
	userID := c.Param("userId")

	// Verify user is admin
	var user models.User
	if err := h.db.First(&user, "id = ? AND role = ?", userID, "admin").Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can view all pending bookings"})
		return
	}

	// Get all pending bookings across all services
	var bookings []models.Booking
	if err := h.db.
		Preload("User").
		Preload("Service").
		Where("status = ?", "pending").
		Order("created_at DESC").
		Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// GetAllBookings returns all bookings across all services (for admin)
func (h *ApprovalHandler) GetAllBookings(c *gin.Context) {
	userID := c.Param("userId")

	// Verify user is admin
	var user models.User
	if err := h.db.First(&user, "id = ? AND role = ?", userID, "admin").Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can view all bookings"})
		return
	}

	// Get all bookings across all services
	var bookings []models.Booking
	if err := h.db.
		Preload("User").
		Preload("Service").
		Order("created_at DESC").
		Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// AdminApproveBooking approves a booking (admin can approve any booking)
func (h *ApprovalHandler) AdminApproveBooking(c *gin.Context) {
	bookingID := c.Param("id")
	userID := c.Param("userId")

	// Verify user is admin
	var admin models.User
	if err := h.db.First(&admin, "id = ? AND role = ?", userID, "admin").Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can approve bookings"})
		return
	}

	var req ApprovalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the booking
	var booking models.Booking
	if err := h.db.Preload("Service").First(&booking, "id = ?", bookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Update booking
	updates := map[string]interface{}{
		"status":         "approved",
		"approved_by":    admin.ID,
		"approval_notes": req.ApprovalNotes,
	}

	if err := h.db.Model(&booking).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve booking"})
		return
	}

	// Reload booking with updated data
	h.db.Preload("Service").First(&booking, "id = ?", bookingID)

	// Create notification for student
	notification := models.Notification{
		UserID:  booking.UserID,
		Title:   "Booking Approved",
		Message: "Your booking for " + booking.Service.Name + " has been approved by admin.",
		Type:    "booking_approval",
		IsRead:  false,
	}
	h.db.Create(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Booking approved successfully", "booking": booking})
}

// AdminRejectBooking rejects a booking (admin can reject any booking)
func (h *ApprovalHandler) AdminRejectBooking(c *gin.Context) {
	bookingID := c.Param("id")
	userID := c.Param("userId")

	// Verify user is admin
	var admin models.User
	if err := h.db.First(&admin, "id = ? AND role = ?", userID, "admin").Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can reject bookings"})
		return
	}

	var req ApprovalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the booking
	var booking models.Booking
	if err := h.db.Preload("Service").First(&booking, "id = ?", bookingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Update booking
	updates := map[string]interface{}{
		"status":         "rejected",
		"approved_by":    admin.ID,
		"approval_notes": req.ApprovalNotes,
	}

	if err := h.db.Model(&booking).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject booking"})
		return
	}

	// Reload booking with updated data
	h.db.Preload("Service").First(&booking, "id = ?", bookingID)

	// Create notification for student
	notification := models.Notification{
		UserID:  booking.UserID,
		Title:   "Booking Rejected",
		Message: "Your booking for " + booking.Service.Name + " has been rejected by admin. Reason: " + req.ApprovalNotes,
		Type:    "booking_rejection",
		IsRead:  false,
	}
	h.db.Create(&notification)

	c.JSON(http.StatusOK, gin.H{"message": "Booking rejected successfully", "booking": booking})
}
