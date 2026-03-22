package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func TestListServicesReturnsAllServices(t *testing.T) {
	db := setupTestDB(t)
	createServiceFixture(t, db)
	createServiceFixture(t, db, func(service *models.Service) {
		service.Name = "Dining Hall"
		service.Category = "dining"
	})

	handler := NewServiceHandler(db)
	router := gin.New()
	router.GET("/services", handler.ListServices)

	rec := performRequest(t, router, http.MethodGet, "/services", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	services := decodeJSON[[]models.Service](t, rec)
	if len(services) != 2 {
		t.Fatalf("expected 2 services, got %d", len(services))
	}
}

func TestGetServiceReturnsNotFoundForMissingService(t *testing.T) {
	db := setupTestDB(t)
	handler := NewServiceHandler(db)
	router := gin.New()
	router.GET("/services/:id", handler.GetService)

	rec := performRequest(t, router, http.MethodGet, "/services/missing", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestCreateServicePersistsRecord(t *testing.T) {
	db := setupTestDB(t)
	handler := NewServiceHandler(db)
	router := gin.New()
	router.POST("/services", handler.CreateService)

	rec := performRequest(t, router, http.MethodPost, "/services", CreateServiceRequest{
		Name:        "Health Center",
		Description: "Clinic support",
		Category:    "health",
		Location:    "Building A",
		Email:       "health@campus.edu",
		Hours:       "Weekdays",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	var service models.Service
	if err := db.Where("name = ?", "Health Center").First(&service).Error; err != nil {
		t.Fatalf("expected service to be persisted: %v", err)
	}
}

func TestUpdateServicePersistsChanges(t *testing.T) {
	db := setupTestDB(t)
	service := createServiceFixture(t, db)

	handler := NewServiceHandler(db)
	router := gin.New()
	router.PUT("/services/:id", handler.UpdateService)

	rec := performRequest(t, router, http.MethodPut, "/services/"+service.ID, CreateServiceRequest{
		Name:        "Updated Shuttle",
		Description: "Updated description",
		Category:    "transportation",
		Location:    "West Hub",
		Email:       "updated@campus.edu",
		Hours:       "Weekdays",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Service
	if err := db.First(&updated, "id = ?", service.ID).Error; err != nil {
		t.Fatalf("failed to reload service: %v", err)
	}

	if updated.Name != "Updated Shuttle" {
		t.Fatalf("expected service name to be updated, got %s", updated.Name)
	}
}

func TestDeleteServiceRemovesRecord(t *testing.T) {
	db := setupTestDB(t)
	service := createServiceFixture(t, db)

	handler := NewServiceHandler(db)
	router := gin.New()
	router.DELETE("/services/:id", handler.DeleteService)

	rec := performRequest(t, router, http.MethodDelete, "/services/"+service.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var count int64
	if err := db.Model(&models.Service{}).Where("id = ?", service.ID).Count(&count).Error; err != nil {
		t.Fatalf("failed to count services: %v", err)
	}
	if count != 0 {
		t.Fatalf("expected service to be deleted, found %d records", count)
	}
}

func TestGetServicesByCategoryFiltersResults(t *testing.T) {
	db := setupTestDB(t)
	createServiceFixture(t, db, func(service *models.Service) {
		service.Category = "health"
	})
	createServiceFixture(t, db, func(service *models.Service) {
		service.Name = "Dining Hall"
		service.Category = "dining"
	})

	handler := NewServiceHandler(db)
	router := gin.New()
	router.GET("/services/category/:category", handler.GetServicesByCategory)

	rec := performRequest(t, router, http.MethodGet, "/services/category/health", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	services := decodeJSON[[]models.Service](t, rec)
	if len(services) != 1 || services[0].Category != "health" {
		t.Fatalf("expected one health service, got %+v", services)
	}
}
