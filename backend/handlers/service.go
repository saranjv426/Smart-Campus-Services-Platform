package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"smart-campus-services/models"
)

type ServiceHandler struct {
	db *gorm.DB
}

func NewServiceHandler(db *gorm.DB) *ServiceHandler {
	return &ServiceHandler{db: db}
}

type CreateServiceRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageURL    string `json:"imageUrl"`
	Location    string `json:"location" binding:"required"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Hours       string `json:"hours"`
}

type UpdateServiceActiveRequest struct {
	IsActive *bool `json:"isActive" binding:"required"`
}

const (
	defaultPageSize = 100
	maxPageSize     = 200
)

var allowedSortColumns = map[string]string{
	"name":      "name",
	"category":  "category",
	"rating":    "rating",
	"createdAt": "created_at",
	"updatedAt": "updated_at",
}

// ListServices returns services with optional filtering, search, and pagination.
//
// Query params:
// - q: case-insensitive partial search on name and description
// - category: exact category match (case-insensitive)
// - activeOnly: true/false, defaults to false
// - limit: max rows to return (1-200), defaults to 100
// - offset: rows to skip, defaults to 0
// - sortBy: one of name, category, rating, createdAt, updatedAt
// - sortOrder: asc/desc, defaults to asc
func (h *ServiceHandler) ListServices(c *gin.Context) {
	query := h.db.Model(&models.Service{})

	searchQuery := strings.TrimSpace(c.Query("q"))
	if searchQuery != "" {
		searchPattern := "%" + strings.ToLower(searchQuery) + "%"
		query = query.Where(
			"LOWER(name) LIKE ? OR LOWER(description) LIKE ?",
			searchPattern,
			searchPattern,
		)
	}

	categoryFilter := strings.TrimSpace(c.Query("category"))
	if categoryFilter != "" {
		query = query.Where("LOWER(category) = ?", strings.ToLower(categoryFilter))
	}

	activeOnly := strings.EqualFold(c.Query("activeOnly"), "true")
	if activeOnly {
		query = query.Where("is_active = ?", true)
	}

	limit := parseBoundedInt(c.Query("limit"), defaultPageSize, 1, maxPageSize)
	offset := parseBoundedInt(c.Query("offset"), 0, 0, 1000000)

	sortColumn := resolveSortColumn(c.Query("sortBy"))
	sortDirection := resolveSortDirection(c.Query("sortOrder"))
	query = query.Order(sortColumn + " " + sortDirection)
	query = query.Limit(limit).Offset(offset)

	var services []models.Service
	if err := query.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch services"})
		return
	}

	c.JSON(http.StatusOK, services)
}

func parseBoundedInt(rawValue string, defaultValue, minValue, maxValue int) int {
	parsedValue, err := strconv.Atoi(rawValue)
	if err != nil {
		return defaultValue
	}
	if parsedValue < minValue {
		return minValue
	}
	if parsedValue > maxValue {
		return maxValue
	}
	return parsedValue
}

func resolveSortColumn(rawColumn string) string {
	if column, ok := allowedSortColumns[strings.TrimSpace(rawColumn)]; ok {
		return column
	}
	return "name"
}

func resolveSortDirection(rawDirection string) string {
	if strings.EqualFold(strings.TrimSpace(rawDirection), "desc") {
		return "DESC"
	}
	return "ASC"
}

// GetService returns a service by ID
func (h *ServiceHandler) GetService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service

	if err := h.db.Preload("Reviews").First(&service, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// CreateService creates a new service
func (h *ServiceHandler) CreateService(c *gin.Context) {
	var req CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := models.Service{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
		Location:    req.Location,
		Phone:       req.Phone,
		Email:       req.Email,
		Hours:       req.Hours,
		IsActive:    true,
	}

	if err := h.db.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}

	c.JSON(http.StatusCreated, service)
}

// UpdateService updates a service
func (h *ServiceHandler) UpdateService(c *gin.Context) {
	id := c.Param("id")
	var req CreateServiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var service models.Service
	if err := h.db.First(&service, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service"})
		return
	}

	if err := h.db.Model(&service).Updates(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// UpdateServiceActive toggles a service's active state without updating other fields.
func (h *ServiceHandler) UpdateServiceActive(c *gin.Context) {
	id := c.Param("id")
	var req UpdateServiceActiveRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.IsActive == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "isActive is required"})
		return
	}

	var service models.Service
	if err := h.db.First(&service, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service"})
		return
	}

	service.IsActive = *req.IsActive
	if err := h.db.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// DeleteService deletes a service
func (h *ServiceHandler) DeleteService(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.Delete(&models.Service{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}

// GetServicesByCategory returns services by category
func (h *ServiceHandler) GetServicesByCategory(c *gin.Context) {
	category := c.Param("category")
	var services []models.Service

	if err := h.db.Where("category = ?", category).Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch services"})
		return
	}

	c.JSON(http.StatusOK, services)
}
