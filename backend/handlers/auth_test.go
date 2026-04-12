package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/middleware"
	"smart-campus-services/models"
	"smart-campus-services/validation"
)

func setupAuthHandlerTest(t *testing.T) (*gin.Engine, *AuthHandler) {
	t.Helper()

	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	return gin.New(), handler
}

func TestRegisterCreatesUser(t *testing.T) {
	router, handler := setupAuthHandlerTest(t)
	router.POST("/register", handler.Register)

	rec := performRequest(t, router, http.MethodPost, "/register", RegisterRequest{
		Email:     "student@campus.edu",
		Password:  "secret123",
		FirstName: "Sam",
		LastName:  "Student",
		Phone:     "555-1000",
		Role:      "student",
	})

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[AuthResponse](t, rec)
	if resp.Email != "student@campus.edu" || resp.Role != "student" || resp.Token == "" {
		t.Fatalf("expected auth response with created user details, got %+v", resp)
	}
}

func TestRegisterRejectsDuplicateEmail(t *testing.T) {
	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := setupTestDB(t)
	createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
	})

	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/register", handler.Register)

	rec := performRequest(t, router, http.MethodPost, "/register", RegisterRequest{
		Email:     "student@campus.edu",
		Password:  "secret123",
		FirstName: "Sam",
		LastName:  "Student",
		Phone:     "555-1000",
		Role:      "student",
	})

	if rec.Code != http.StatusConflict {
		t.Fatalf("expected status 409, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLoginAuthenticatesUser(t *testing.T) {
	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := setupTestDB(t)
	user := createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/login", handler.Login)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "secret123",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[AuthResponse](t, rec)
	if resp.ID != user.ID || resp.Email != user.Email || resp.Token == "" {
		t.Fatalf("expected authenticated user in response, got %+v", resp)
	}
}

func TestLoginRejectsInvalidPassword(t *testing.T) {
	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := setupTestDB(t)
	createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := gin.New()
	router.POST("/login", handler.Login)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "wrongpass",
	})

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLogoutReturnsSuccess(t *testing.T) {
	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := setupTestDB(t)
	handler := NewAuthHandler(db)

	router := gin.New()
	router.POST("/logout", handler.Logout)

	rec := performRequest(t, router, http.MethodPost, "/logout", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestRefreshTokenReturnsStructuredToken(t *testing.T) {
	if err := validation.Init(); err != nil {
		t.Fatalf("failed to initialize validator: %v", err)
	}

	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	user := createUserFixture(t, db, func(user *models.User) {
		user.Role = "admin"
	})

	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set("userID", user.ID)
		c.Set("role", user.Role)
		c.Set("serviceID", user.ServiceID)
		c.Next()
	})
	router.POST("/refresh", handler.RefreshToken)

	rec := performRequest(t, router, http.MethodPost, "/refresh", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	body := decodeJSON[map[string]string](t, rec)
	claims, ok := middleware.ParseTokenClaims(body["token"])
	if !ok {
		t.Fatalf("expected refresh token to be parseable, got %q", body["token"])
	}
	if claims.UserID != user.ID || claims.Role != user.Role || claims.ServiceID != user.ServiceID {
		t.Fatalf("unexpected refresh token claims: %+v", claims)
	}
}
