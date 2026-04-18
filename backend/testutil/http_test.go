package testutil

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMustMarshalReturnsJSONBytes(t *testing.T) {
	body := MustMarshal(map[string]string{"hello": "world"})

	if string(body) != "{\"hello\":\"world\"}" {
		t.Fatalf("expected valid JSON bytes, got %s", string(body))
	}
}

func TestPerformRequestSetsJSONContentTypeWhenBodyPresent(t *testing.T) {
	router := gin.New()
	router.POST("/echo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"contentType": c.GetHeader("Content-Type")})
	})

	rec := PerformRequest(router, http.MethodPost, "/echo", MustMarshal(map[string]string{"x": "y"}))

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d body=%s", rec.Code, rec.Body.String())
	}
	if rec.Body.String() != "{\"contentType\":\"application/json\"}" {
		t.Fatalf("expected JSON content type to be set, got %s", rec.Body.String())
	}
}

func TestNewJSONRequestLeavesContentTypeEmptyWithoutBody(t *testing.T) {
	req := NewJSONRequest(t, http.MethodGet, "/health", nil)

	if req.Header.Get("Content-Type") != "" {
		t.Fatalf("expected no content type for empty body, got %q", req.Header.Get("Content-Type"))
	}
}
