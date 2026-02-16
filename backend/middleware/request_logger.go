package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestIDHeader = "X-Request-ID"

// RequestLogger adds request IDs and logs request latency + status.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(requestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Set("requestID", requestID)
		c.Writer.Header().Set(requestIDHeader, requestID)

		start := time.Now()
		c.Next()

		log.Printf("request_id=%s method=%s path=%s status=%d duration=%s",
			requestID,
			c.Request.Method,
			c.FullPath(),
			c.Writer.Status(),
			time.Since(start),
		)
	}
}
