package handlers

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"smart-campus-services/models"
)

type BookingHandler struct {
	db *gorm.DB
}

var allowedBookingStatuses = map[string]struct{}{
	"pending":   {},
	"approved":  {},
	"rejected":  {},
	"completed": {},
	"cancelled": {},
}

func NewBookingHandler(db *gorm.DB) *BookingHandler {
	return &BookingHandler{db: db}
}

type CreateBookingRequest struct {
	UserID    string    `json:"userId" binding:"required"`
	ServiceID string    `json:"serviceId" binding:"required"`
	StartTime time.Time `json:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" binding:"required"`
	Notes     string    `json:"notes"`
}

// CreateBooking creates a new booking
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !req.EndTime.After(req.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endTime must be after startTime"})
		return
	}

	booking := models.Booking{
		UserID:    req.UserID,
		ServiceID: req.ServiceID,
		Status:    "pending",
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Notes:     req.Notes,
	}

	if err := h.db.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// GetBooking returns a booking by ID
func (h *BookingHandler) GetBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking

	if err := h.db.Preload("User").Preload("Service").First(&booking, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// GetAllBookings returns all bookings
func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	var bookings []models.Booking

	if err := h.db.Preload("User").Preload("Service").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// GetUserBookings returns all bookings for a user
func (h *BookingHandler) GetUserBookings(c *gin.Context) {
	userId := c.Param("userId")
	status := normalizeBookingStatus(c.Query("status"))
	var bookings []models.Booking

	query := h.db.Preload("Service").Where("user_id = ?", userId)

	if status != "" {
		if !isAllowedBookingStatus(status) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid status. Supported statuses: pending, approved, rejected, completed, cancelled",
			})
			return
		}

		query = query.Where("status = ?", status)
	}

	if err := query.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func normalizeBookingStatus(status string) string {
	return strings.ToLower(strings.TrimSpace(status))
}

func isAllowedBookingStatus(status string) bool {
	_, ok := allowedBookingStatuses[status]
	return ok
}

// UpdateBooking updates a booking
func (h *BookingHandler) UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var req CreateBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !req.EndTime.After(req.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "endTime must be after startTime"})
		return
	}

	var booking models.Booking
	if err := h.db.First(&booking, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		return
	}

	if err := h.db.Model(&booking).Where("id = ?", id).Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}
	if err := h.db.First(&booking, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated booking"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// CancelBooking cancels a booking
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	id := c.Param("id")
	result := h.db.Model(&models.Booking{}).Where("id = ?", id).Update("status", "cancelled")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}
