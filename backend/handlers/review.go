package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"smart-campus-services/models"
)

type ReviewHandler struct {
	db *gorm.DB
}

func NewReviewHandler(db *gorm.DB) *ReviewHandler {
	return &ReviewHandler{db: db}
}

type CreateReviewRequest struct {
	UserID    string `json:"userId" binding:"required"`
	ServiceID string `json:"serviceId" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	Comment   string `json:"comment" binding:"required"`
}

// CreateReview creates a new review
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Review{
		UserID:    req.UserID,
		ServiceID: req.ServiceID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	}

	if err := h.db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	// Update service rating
	h.updateServiceRating(req.ServiceID)

	c.JSON(http.StatusCreated, review)
}

// GetServiceReviews returns all reviews for a service
func (h *ReviewHandler) GetServiceReviews(c *gin.Context) {
	serviceId := c.Param("serviceId")
	var reviews []models.Review

	if err := h.db.Preload("User").Where("service_id = ?", serviceId).Order("created_at DESC").Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReview returns a review by ID
func (h *ReviewHandler) GetReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := h.db.Preload("User").Preload("Service").First(&review, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch review"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// DeleteReview deletes a review
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := h.db.First(&review, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	if err := h.db.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	// Update service rating
	h.updateServiceRating(review.ServiceID)

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

// updateServiceRating updates the service rating based on all reviews
func (h *ReviewHandler) updateServiceRating(serviceId string) {
	var avgRating float64
	h.db.Model(&models.Review{}).Where("service_id = ?", serviceId).Select("COALESCE(AVG(rating), 0)").Row().Scan(&avgRating)

	h.db.Model(&models.Service{}).Where("id = ?", serviceId).Update("rating", avgRating)
}
