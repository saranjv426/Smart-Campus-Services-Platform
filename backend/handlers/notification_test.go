package handlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func TestGetNotificationsReturnsUserNotifications(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	otherUser := createUserFixture(t, db)
	createNotificationFixture(t, db, user.ID)
	createNotificationFixture(t, db, otherUser.ID)

	handler := NewNotificationHandler(db)
	router := gin.New()
	router.GET("/notifications/:userId", handler.GetNotifications)

	rec := performRequest(t, router, http.MethodGet, "/notifications/"+user.ID, nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	notifications := decodeJSON[[]models.Notification](t, rec)
	if len(notifications) != 1 || notifications[0].UserID != user.ID {
		t.Fatalf("expected one notification for user %s, got %+v", user.ID, notifications)
	}
}

func TestCreateNotificationDefaultsUnread(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)

	handler := NewNotificationHandler(db)
	router := gin.New()
	router.POST("/notifications", handler.CreateNotification)

	rec := performRequest(t, router, http.MethodPost, "/notifications", CreateNotificationRequest{
		UserID:  user.ID,
		Title:   "Approval Update",
		Message: "Your booking was reviewed",
		Type:    "booking",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	notification := decodeJSON[models.Notification](t, rec)
	if notification.IsRead {
		t.Fatal("expected notification to be unread by default")
	}
}

func TestMarkAsReadUpdatesNotification(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db)
	notification := createNotificationFixture(t, db, user.ID, func(n *models.Notification) {
		n.CreatedAt = time.Now().Add(-time.Hour)
	})

	handler := NewNotificationHandler(db)
	router := gin.New()
	router.PUT("/notifications/:id/read", handler.MarkAsRead)

	rec := performRequest(t, router, http.MethodPut, "/notifications/"+notification.ID+"/read", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	var updated models.Notification
	if err := db.First(&updated, "id = ?", notification.ID).Error; err != nil {
		t.Fatalf("failed to reload notification: %v", err)
	}
	if !updated.IsRead {
		t.Fatal("expected notification to be marked as read")
	}
}

func TestMarkAsReadReturnsNotFoundForMissingNotification(t *testing.T) {
	db := setupTestDB(t)
	handler := NewNotificationHandler(db)
	router := gin.New()
	router.PUT("/notifications/:id/read", handler.MarkAsRead)

	rec := performRequest(t, router, http.MethodPut, "/notifications/missing/read", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d with body %s", rec.Code, rec.Body.String())
	}
}
