package handlers

import (
	"encoding/json"
	"net/http"
	"testing"

	"smart-campus-services/models"
	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
)

func setupNotificationRouter(t *testing.T) (*gin.Engine, string, string) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	h := NewNotificationHandler(db)

	user := models.User{
		Email:     "notif@campus.edu",
		Password:  "password123",
		FirstName: "Notif",
		LastName:  "User",
		Phone:     "+2222222222",
		Role:      "student",
	}
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("failed to seed user: %v", err)
	}

	notification := models.Notification{
		UserID:  user.ID,
		Title:   "Welcome",
		Message: "Welcome to campus services",
		Type:    "announcement",
		IsRead:  false,
	}
	if err := db.Create(&notification).Error; err != nil {
		t.Fatalf("failed to seed notification: %v", err)
	}

	r := gin.New()
	r.GET("/api/notifications/:userId", h.GetNotifications)
	r.POST("/api/notifications", h.CreateNotification)
	r.PUT("/api/notifications/:id/read", h.MarkAsRead)

	return r, user.ID, notification.ID
}

func TestGetNotificationsSuccess(t *testing.T) {
	r, userID, _ := setupNotificationRouter(t)

	resp := testutil.PerformRequest(r, http.MethodGet, "/api/notifications/"+userID, nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	var notifications []map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &notifications); err != nil {
		t.Fatalf("failed to decode notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
}

func TestCreateNotificationValidation(t *testing.T) {
	r, userID, _ := setupNotificationRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"userId": userID,
		"title":  "Incomplete payload",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/notifications", body)
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}

func TestCreateNotificationSuccess(t *testing.T) {
	r, userID, _ := setupNotificationRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"userId":  userID,
		"title":   "Booking Update",
		"message": "Your booking is confirmed",
		"type":    "booking",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/notifications", body)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d body=%s", http.StatusCreated, resp.Code, resp.Body.String())
	}

	var created map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to decode notification: %v", err)
	}
	if created["isRead"] != false {
		t.Fatalf("expected created notification to be unread, got %v", created["isRead"])
	}
}

func TestGetNotificationsEmptyList(t *testing.T) {
	r, _, _ := setupNotificationRouter(t)

	resp := testutil.PerformRequest(r, http.MethodGet, "/api/notifications/non-existent-user", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	var notifications []map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &notifications); err != nil {
		t.Fatalf("failed to decode notifications: %v", err)
	}
	if len(notifications) != 0 {
		t.Fatalf("expected 0 notifications, got %d", len(notifications))
	}
}

func TestMarkAsReadSuccess(t *testing.T) {
	r, userID, notificationID := setupNotificationRouter(t)

	resp := testutil.PerformRequest(r, http.MethodPut, "/api/notifications/"+notificationID+"/read", nil)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, resp.Code, resp.Body.String())
	}

	listResp := testutil.PerformRequest(r, http.MethodGet, "/api/notifications/"+userID, nil)
	if listResp.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d body=%s", http.StatusOK, listResp.Code, listResp.Body.String())
	}

	var notifications []map[string]any
	if err := json.Unmarshal(listResp.Body.Bytes(), &notifications); err != nil {
		t.Fatalf("failed to decode notifications: %v", err)
	}
	if len(notifications) != 1 {
		t.Fatalf("expected 1 notification, got %d", len(notifications))
	}
	if notifications[0]["isRead"] != true {
		t.Fatalf("expected notification to be read, got %v", notifications[0]["isRead"])
	}
}
