package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func TestCreateReviewPersistsRecordAndUpdatesServiceRating(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)

	handler := NewReviewHandler(db)
	router := gin.New()
	router.POST("/reviews", handler.CreateReview)

	rec := performRequest(t, router, http.MethodPost, "/reviews", CreateReviewRequest{
		UserID:    user.ID,
		ServiceID: service.ID,
		Rating:    4,
		Comment:   "Very helpful",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Service
	if err := db.First(&updated, "id = ?", service.ID).Error; err != nil {
		t.Fatalf("failed to reload service: %v", err)
	}
	if updated.Rating != 4 {
		t.Fatalf("expected service rating 4, got %v", updated.Rating)
	}
}

func TestGetServiceReviewsReturnsReviewsForService(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	otherService := createServiceFixture(t, db, func(s *models.Service) {
		s.Name = "Dining Hall"
		s.Category = "dining"
	})
	createReviewFixture(t, db, user.ID, service.ID, 5)
	createReviewFixture(t, db, user.ID, otherService.ID, 3)

	handler := NewReviewHandler(db)
	router := gin.New()
	router.GET("/reviews/service/:serviceId", handler.GetServiceReviews)

	rec := performRequest(t, router, http.MethodGet, "/reviews/service/"+service.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	reviews := decodeJSON[[]models.Review](t, rec)
	if len(reviews) != 1 || reviews[0].ServiceID != service.ID {
		t.Fatalf("expected one review for service %s, got %+v", service.ID, reviews)
	}
}

func TestGetReviewReturnsSingleReview(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	review := createReviewFixture(t, db, user.ID, service.ID, 5)

	handler := NewReviewHandler(db)
	router := gin.New()
	router.GET("/reviews/:id", handler.GetReview)

	rec := performRequest(t, router, http.MethodGet, "/reviews/"+review.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestGetUserReviewsReturnsReviewsForUser(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	otherUser := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	createReviewFixture(t, db, user.ID, service.ID, 4)
	createReviewFixture(t, db, otherUser.ID, service.ID, 2)

	handler := NewReviewHandler(db)
	router := gin.New()
	router.GET("/reviews/user/:userId", handler.GetUserReviews)

	rec := performRequest(t, router, http.MethodGet, "/reviews/user/"+user.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	reviews := decodeJSON[[]models.Review](t, rec)
	if len(reviews) != 1 || reviews[0].UserID != user.ID {
		t.Fatalf("expected one review for user %s, got %+v", user.ID, reviews)
	}
}

func TestDeleteReviewRespectsOwnershipAndUpdatesRating(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	service := createServiceFixture(t, db)
	otherReview := createReviewFixture(t, db, user.ID, service.ID, 5)
	review := createReviewFixture(t, db, user.ID, service.ID, 3, func(r *models.Review) {
		r.Comment = "Delete me"
	})

	handler := NewReviewHandler(db)
	router := gin.New()
	router.DELETE("/reviews/:id", func(c *gin.Context) {
		c.Set("userId", user.ID)
		handler.DeleteReview(c)
	})

	rec := performRequest(t, router, http.MethodDelete, "/reviews/"+review.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var count int64
	if err := db.Model(&models.Review{}).Where("id = ?", review.ID).Count(&count).Error; err != nil {
		t.Fatalf("failed to count reviews: %v", err)
	}
	if count != 0 {
		t.Fatalf("expected review to be deleted, found %d records", count)
	}

	var updated models.Service
	if err := db.First(&updated, "id = ?", service.ID).Error; err != nil {
		t.Fatalf("failed to reload service: %v", err)
	}
	if updated.Rating != float64(otherReview.Rating) {
		t.Fatalf("expected rating to be recalculated to %d, got %v", otherReview.Rating, updated.Rating)
	}
}
