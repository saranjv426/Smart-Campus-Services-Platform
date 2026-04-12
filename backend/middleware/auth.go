package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TokenClaims struct {
	UserID    string
	Role      string
	ServiceID string
}

// ParseTokenClaims extracts lightweight user claims from the demo auth token.
func ParseTokenClaims(token string) (TokenClaims, bool) {
	parts := strings.Split(token, "|")
	if len(parts) != 3 {
		return TokenClaims{}, false
	}

	claims := TokenClaims{
		UserID:    strings.TrimSpace(parts[0]),
		Role:      strings.TrimSpace(parts[1]),
		ServiceID: strings.TrimSpace(parts[2]),
	}
	if claims.UserID == "" || claims.Role == "" {
		return TokenClaims{}, false
	}

	return claims, true
}

// AuthRequired checks for a Bearer token header.
// Replace token parsing with JWT verification in production.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid authorization header"})
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		c.Set("token", token)
		if claims, ok := ParseTokenClaims(token); ok {
			c.Set("userID", claims.UserID)
			c.Set("role", claims.Role)
			c.Set("serviceID", claims.ServiceID)
			if claims.Role == "staff" {
				c.Set("staffID", claims.UserID)
			}
		}
		c.Next()
	}
}

// RequireRoles enforces role-based access using role set in context.
func RequireRoles(roles ...string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		allowed[role] = struct{}{}
	}

	return func(c *gin.Context) {
		role := c.GetString("role")
		if _, ok := allowed[role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}
