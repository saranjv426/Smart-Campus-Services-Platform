package handlers

import (
	"net/http"
	"testing"
	"time"

	"smart-campus-services/testutil"

	"github.com/gin-gonic/gin"
)

func setupBookingRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db := testutil.NewTestDB(t)
	h := NewBookingHandler(db)

	r := gin.New()
	r.POST("/api/bookings", h.CreateBooking)
	return r
}

func TestCreateBookingValidation(t *testing.T) {
	r := setupBookingRouter(t)

	body := testutil.MustMarshal(map[string]any{
		"userId":    "user-1",
		"serviceId": "service-1",
		"startTime": time.Now().UTC(),
		"notes":     "Need slot",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/bookings", body)
	if resp.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusBadRequest, resp.Code, resp.Body.String())
	}
}

func TestCreateBookingSuccess(t *testing.T) {
	r := setupBookingRouter(t)
	start := time.Now().UTC().Add(2 * time.Hour)
	end := start.Add(time.Hour)

	body := testutil.MustMarshal(map[string]any{
		"userId":    "user-1",
		"serviceId": "service-1",
		"startTime": start,
		"endTime":   end,
		"notes":     "Need slot",
	})

	resp := testutil.PerformRequest(r, http.MethodPost, "/api/bookings", body)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d body=%s", http.StatusCreated, resp.Code, resp.Body.String())
	}
}
