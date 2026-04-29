package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"smart-campus-services/models"
)

func setupAuthTestRouter(dbHandler *AuthHandler) *gin.Engine {
	router := gin.New()
	router.POST("/register", dbHandler.Register)
	router.POST("/login", dbHandler.Login)
	router.POST("/logout", dbHandler.Logout)
	router.POST("/refresh", dbHandler.RefreshToken)
	return router
}

func TestRegisterCreatesUser(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

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

	var user models.User
	if err := db.Where("email = ?", "student@campus.edu").First(&user).Error; err != nil {
		t.Fatalf("expected user to be persisted: %v", err)
	}
}

func TestRegisterRejectsDuplicateEmail(t *testing.T) {
	db := setupTestDB(t)
	createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
	})

	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

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

func TestRegisterRejectsInvalidRole(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/register", map[string]any{
		"email":     "student@campus.edu",
		"password":  "secret123",
		"firstName": "Sam",
		"lastName":  "Student",
		"phone":     "555-1000",
		"role":      "visitor",
	})

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLoginAuthenticatesUser(t *testing.T) {
	db := setupTestDB(t)
	user := createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "secret123",
	})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}

	resp := decodeJSON[AuthResponse](t, rec)
	if resp.ID != user.ID {
		t.Fatalf("expected response ID %s, got %s", user.ID, resp.ID)
	}
}

func TestLoginRejectsInvalidPassword(t *testing.T) {
	db := setupTestDB(t)
	createUserFixture(t, db, func(user *models.User) {
		user.Email = "student@campus.edu"
		user.Password = "secret123"
	})

	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/login", LoginRequest{
		Email:    "student@campus.edu",
		Password: "wrongpass",
	})

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestLogoutReturnsSuccess(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/logout", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}

func TestRefreshTokenReturnsPlaceholderToken(t *testing.T) {
	db := setupTestDB(t)
	handler := NewAuthHandler(db)
	router := setupAuthTestRouter(handler)

	rec := performRequest(t, router, http.MethodPost, "/refresh", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", rec.Code, rec.Body.String())
	}
}
